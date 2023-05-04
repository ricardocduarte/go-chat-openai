package interfaces

import (
	"chat-openai/models"
	"context"
)

// ClientInterface is the interface for the OpenAI API client.
// Documentation at https://platform.openai.com/docs/api-reference/introduction
type ClientInterface interface {
	// ListModels Lists the currently available models, and provides basic information about each one such as the owner and availability.
	ListModels(ctx context.Context) (response models.ListModelsResponse, err error)

	// RetrieveModel Retrieves a model instance, providing basic information about the model such as the owner and permissioning.
	RetrieveModel(ctx context.Context, id string) (response models.RetrieveModelResponse, err error)

	// CreateCompletions Creates a completion for the provided prompt and parameters
	CreateCompletions(ctx context.Context, r models.CreateCompletionsRequest) (response models.CreateCompletionsResponse, err error)

	// CreateChats Creates a completion for the chat message
	CreateChats(ctx context.Context, r models.CreateChatsRequest) (response models.CreateChatsResponse, err error)

	// CreateEdits Creates a new edit for the provided input, instruction, and parameters.
	CreateEdits(ctx context.Context, r models.CreateEditsRequest) (response models.CreateEditsResponse, err error)

	// CreateImages Creates an image given a prompt.
	CreateImages(ctx context.Context, r models.CreateImagesRequest) (response models.CreateImagesResponse, err error)

	// CreateImagesEdits Creates an edited or extended image given an original image and a prompt.
	CreateImagesEdits(ctx context.Context, r models.CreateImagesEditsRequest) (response models.CreateImagesEditsResponse, err error)

	// CreateImagesVariations Creates a variation of a given image.
	CreateImagesVariations(ctx context.Context, r models.CreateImagesVariationsRequest) (response models.CreateImagesVariationsResponse, err error)

	// CreateEmbeddings Creates an embedding vector representing the input text.
	CreateEmbeddings(ctx context.Context, r models.CreateEmbeddingsRequest) (response models.CreateEmbeddingsResponse, err error)

	// CreateTranscriptions Transcribes audio into the input language.
	CreateTranscriptions(ctx context.Context, r models.CreateTranscriptionsRequest) (response models.CreateTranscriptionsResponse, err error)

	// CreateTranslations Translates audio into into English.
	CreateTranslations(ctx context.Context, r models.CreateTranslationsRequest) (response models.CreateTranslationsResponse, err error)

	// CreateModerations Classifies if text violates OpenAI's Content Policy
	CreateModerations(ctx context.Context, r models.CreateModerationsRequest) (response models.CreateModerationsResponse, err error)

	// Not yet implemented:
	// ListFiles(ctx context.Context) (response ListFilesResponse, err error)
	// UploadFile(ctx context.Context, r UploadFileRequest) (response UploadFileResponse, err error)
	// DeleteFile(ctx context.Context, id string) (response DeleteFileResponse, err error)
	// RetrieveFile(ctx context.Context, id string) (response RetrieveFileResponse, err error)
	// RetrieveFileContent(ctx context.Context, id string) (response RetrieveFileContentResponse, err error)
	// CreateFineTune(ctx context.Context, r CreateFineTuneRequest) (response CreateFineTuneResponse, err error)
	// ListFineTunes(ctx context.Context) (response ListFineTunesResponse, err error)
	// RetrieveFineTunes(ctx context.Context, id string) (response RetrieveFineTunesResponse, err error)
	// ListFineTuneEvents(ctx context.Context, id string) (response ListFineTuneEventsResponse, err error)
	// DeleteFineTuneModel(ctx context.Context, id string) (response DeleteFineTuneModelResponse, err error)
}
