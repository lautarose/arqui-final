package services

import (
	"context"
	"fmt"
	"io"
	"items/clients/queue"
	"items/dtos"
	"items/services/repositories"
	e "items/utils/errors/errors"
	"net/http"
	"os"
	"path/filepath"
)

type ServiceImpl struct {
	localCache repositories.Repository
	distCache  repositories.Repository
	db         repositories.Repository
	queue      queue.Publisher
}

func NewServiceImpl(
	localCache repositories.Repository,
	distCache repositories.Repository,
	db repositories.Repository,
	queue queue.Publisher,
) *ServiceImpl {
	return &ServiceImpl{
		localCache: localCache,
		distCache:  distCache,
		db:         db,
		queue:      queue,
	}
}

func (serv *ServiceImpl) GetItemById(ctx context.Context, id string) (dtos.ItemDto, e.ApiError) {
	var item dtos.ItemDto
	var items dtos.ItemsDto
	var apiErr e.ApiError
	var source string

	items = append(items, item)

	// try to find it in localCache
	item, apiErr = serv.localCache.GetItemById(ctx, id)
	if apiErr != nil {
		if apiErr.Status() != http.StatusNotFound {
			return dtos.ItemDto{}, apiErr
		}
		// try to find it in distCache
		item, apiErr = serv.distCache.GetItemById(ctx, id)
		if apiErr != nil {
			if apiErr.Status() != http.StatusNotFound {
				return dtos.ItemDto{}, apiErr
			}
			// try to find it in db
			item, apiErr = serv.db.GetItemById(ctx, id)
			if apiErr != nil {
				if apiErr.Status() != http.StatusNotFound {
					return dtos.ItemDto{}, apiErr
				} else {
					fmt.Println(fmt.Sprintf("Not found item %s in any datasource", id))
					apiErr = e.NewNotFoundApiError(fmt.Sprintf("item %s not found", id))
					return dtos.ItemDto{}, apiErr
				}
			} else {
				source = "db"
				defer func() {
					if _, apiErr := serv.distCache.InsertItems(ctx, items); apiErr != nil {
						fmt.Println(fmt.Sprintf("Error trying to save item in distCache %v", apiErr))
					}
					if _, apiErr := serv.localCache.InsertItems(ctx, items); apiErr != nil {
						fmt.Println(fmt.Sprintf("Error trying to save item in localCache %v", apiErr))
					}
				}()
			}
		} else {
			source = "distCache"
			defer func() {
				if _, apiErr := serv.localCache.InsertItems(ctx, items); apiErr != nil {
					fmt.Println(fmt.Sprintf("Error trying to save item in localCache %v", apiErr))
				}
			}()
		}
	} else {
		source = "localCache"
	}

	fmt.Println(fmt.Sprintf("Obtained item from %s!", source))
	return item, nil
}

func (serv *ServiceImpl) InsertItems(ctx context.Context, items dtos.ItemsDto) (dtos.ItemsDto, e.ApiError) {
	results, apiErr := serv.db.InsertItems(ctx, items)
	if apiErr != nil {
		fmt.Println(fmt.Sprintf("Error inserting item in db: %v", apiErr))
		return dtos.ItemsDto{}, apiErr
	}
	fmt.Println(fmt.Sprintf("Inserted item in db: %v", results))

	_, apiErr = serv.distCache.InsertItems(ctx, results)
	if apiErr != nil {
		fmt.Println(fmt.Sprintf("Error inserting item in distCache: %v", apiErr))
		return results, nil
	}
	fmt.Println(fmt.Sprintf("Inserted item in distCache: %v", results))

	_, apiErr = serv.localCache.InsertItems(ctx, results)
	if apiErr != nil {
		fmt.Println(fmt.Sprintf("Error inserting item in localCache: %v", apiErr))
		return results, nil
	}
	fmt.Println(fmt.Sprintf("Inserted item in localCache: %v", results))

	for _, item := range results {
		if err := serv.queue.Publish(ctx, item); err != nil {
			return results, e.NewInternalServerApiError(fmt.Sprintf("Error publishing item %s", item.Id), err)
		}
		fmt.Println(fmt.Sprintf("Message sent: %v", item.Id))

		go downloadImage(item.Picture, item.Id, "/home")
	}

	return results, nil
}

func downloadImage(url, name, folder string) error {
	// Crear la ruta completa de la imagen

	name = "image-" + name
	filePath := filepath.Join(folder, name)

	// Hacer una solicitud HTTP GET a la URL de la imagen
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error image request HTTP: %s", err)
	}
	defer response.Body.Close()

	// Crear el archivo en la carpeta de destino
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %s", err)
	}
	defer file.Close()

	// Copiar el contenido de la respuesta HTTP al archivo

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return fmt.Errorf("error saving image: %s", err)
	}

	fmt.Printf("%s downloaded succesfully! \n", name)
	return nil
}
