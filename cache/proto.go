package cache

import "github.com/golang/protobuf/proto"

func GetProtoObject() proto.Message {
	proto := &Photo{
		Id:                 1,
		PhotoId:            "photo_id_1",
		TypeField:          Photo_USER,
		EntityType:         Photo_SOME_TYPE,
		EntityId:           1,
		SomeCategory:       Photo_SOME_PAGE,
		SomeOtherCategory:  Photo_SOME_OTHER_CATEGORY,
		Caption:            "Dummy caption for " + "id_1",
		UserId:             3123841272,
		SomeStatus:         Photo_STATUS,
		Source:             Photo_ANDROID,
		SomethingBy:        1193812319231,
		Timestamp:          1577817000,
		Filename:           "photo_dummy_filename_long.jpg",
		UpdateTimestamp:    1577818000,
		DeletionReason:     Photo_SOME_REASON,
		ActiveStatus:       Photo_ACTIVE,
		Width:              1000,
		Height:             1000,
		Position:           10,
		SomethingTimestamp: 1577819000,
		Metadata: &Metadata{
			SomeId: "1230128312",
			QualityData: &QualityData{
				QualityDataPresent: true,
				QualityScore:       82.5,
			},
			ClassificationData: &ClassificationData{
				ClassificationDataPresent: true,
				ClassificationCategory:    ClassificationData_SOME_CLASSIFICATION_CATEGORY,
				Probability:               0.1,
			},
			DeduplicationData: &DeduplicationData{
				IsDeduplicationDataPresent: true,
				DuplicatePhotoId:           "photo_id_2",
				SimilarityScore:            0.92,
			},
		},
	}

	return proto
}
