package domain

const (
	CropErrorInvalidArea = iota
	CropErrorInvalidCropType
	CropErrorInvalidCropStatus

	// Crop move to area errors
	CropMoveToAreaErrorInvalidSourceArea
	CropMoveToAreaErrorSourceAreaNotFound
	CropMoveToAreaErrorInvalidDestinationArea
	CropMoveToAreaErrorDestinationAreaNotFound
	CropMoveToAreaErrorInvalidQuantity
	CropMoveToAreaErrorInvalidAreaRules
	CropMoveToAreaErrorInvalidExistingSourceArea
	CropMoveToAreaErrorCannotBeSame
	CropMoveToAreaErrorInvalidExistingArea

	// Crop harvest errors
	CropHarvestErrorInvalidSourceArea
	CropHarvestErrorSourceAreaNotFound
	CropHarvestErrorInvalidQuantity
	CropHarvestErrorInvalidHarvestType

	// Crop dump errors
	CropDumpErrorInvalidSourceArea
	CropDumpErrorSourceAreaNotFound
	CropDumpErrorInvalidQuantity

	// Crop water errors
	CropWaterErrorInvalidWateringDate
	CropWaterErrorInvalidSourceArea
	CropWaterErrorSourceAreaNotFound

	// Crop Batch ID errors
	CropErrorInvalidBatchID
	CropErrorBatchIDAlreadyCreated

	CropContainerErrorInvalidType
	CropContainerErrorInvalidQuantity
	CropContainerErrorInvalidTrayCell

	CropMaterialErrorInvalidMaterial
	CropMaterialErrorNotFound

	CropNoteErrorInvalidContent
	CropNoteErrorNotFound
)

// CropError is a custom error from Go built-in error
type CropError struct {
	Code int
}

func (e CropError) Error() string {
	switch e.Code {
	case CropErrorInvalidArea:
		return "Invalid area"
	case CropErrorInvalidCropType:
		return "Invalid crop type"
	case CropErrorInvalidCropStatus:
		return "Invalid crop status"
	case CropErrorInvalidBatchID:
		return "Invalid crop batch ID"
	case CropErrorBatchIDAlreadyCreated:
		return "Crop batch ID already created"

	case CropMoveToAreaErrorInvalidSourceArea:
		return "Crop source area is invalid"
	case CropMoveToAreaErrorSourceAreaNotFound:
		return "Crop source area not found"
	case CropMoveToAreaErrorInvalidDestinationArea:
		return "Crop destination area is invalid"
	case CropMoveToAreaErrorDestinationAreaNotFound:
		return "Crop destination not found"
	case CropMoveToAreaErrorInvalidQuantity:
		return "Invalid quantity. Make sure your quantity is not zero and enough to be moved"
	case CropMoveToAreaErrorInvalidAreaRules:
		return "Invalid move crop to area. Crop can only be moved from Seeding to Growing, Seeding to Seeding or Growing to Growing"
	case CropMoveToAreaErrorInvalidExistingSourceArea:
		return "Invalid existing source area"
	case CropMoveToAreaErrorCannotBeSame:
		return "Invalid move crop to area. Area source and destination cannot be same"
	case CropMoveToAreaErrorInvalidExistingArea:
		return "invalid existing area. Make sure your existing area is there"

	case CropHarvestErrorInvalidSourceArea:
		return "Invalid source area"
	case CropHarvestErrorSourceAreaNotFound:
		return "Source area not found"
	case CropHarvestErrorInvalidQuantity:
		return "Invalid quantity"
	case CropHarvestErrorInvalidHarvestType:
		return "Invalid harvest type"

	case CropDumpErrorInvalidSourceArea:
		return "Invalid source area"
	case CropDumpErrorSourceAreaNotFound:
		return "Source area not found"
	case CropDumpErrorInvalidQuantity:
		return "Invalid quantity"

	case CropWaterErrorInvalidWateringDate:
		return "Invalid watering date"
	case CropWaterErrorInvalidSourceArea:
		return "Invalid source area"
	case CropWaterErrorSourceAreaNotFound:
		return "Source area not found"

	case CropContainerErrorInvalidType:
		return "Invalid crop container type"
	case CropContainerErrorInvalidQuantity:
		return "Invalid crop container quantity"
	case CropContainerErrorInvalidTrayCell:
		return "Invalid crop container tray cell"

	case CropMaterialErrorInvalidMaterial:
		return "Invalid crop material"
	case CropMaterialErrorNotFound:
		return "Crop inventory material not found"

	case CropNoteErrorInvalidContent:
		return "Invalid crop note content"
	case CropNoteErrorNotFound:
		return "Crop note not found"
	default:
		return "Unrecognized Crop Error Code"
	}
}
