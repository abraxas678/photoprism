package entity

import (
	"time"
)

var editTime = time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC)

var PhotoFixtures = map[string]Photo{
	"19800101_000002_D640C559": {
		ID:               1000000,
		PhotoUUID:        "pt9jtdre2lvl0yh7",
		TakenAt:          time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC),
		TakenAtLocal:     time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC),
		TakenSrc:         "exif",
		PhotoTitle:       "",
		TitleSrc:         "",
		PhotoPath:        "2790/02",
		PhotoName:        "19800101_000002_D640C559",
		PhotoQuality:     3,
		PhotoResolution:  2,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoStory:       false,
		PhotoLat:         48.519234,
		PhotoLng:         9.057997,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraID:         0,
		CameraSerial:     "",
		CameraSrc:        "",
		LensID:           0,
		PlaceID:          "zz",
		LocationID:       "",
		LocationSrc:      "",
		TimeZone:         "",
		PhotoCountry:     "zz",
		PhotoYear:        2790,
		PhotoMonth:       2,
		Description:      DescriptionFixtureLake,
		DescriptionSrc:   "",
		Camera:           &CameraFixtureEOS6D,
		Lens:             nil,
		Location:         nil,
		Place:            nil,
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{},
		CreatedAt:        time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
	"Photo01": {
		ID:               1000001,
		PhotoUUID:        "pt9jtdre2lvl0yh8",
		TakenAt:          time.Date(2006, 1, 1, 2, 0, 0, 0, time.UTC),
		TakenAtLocal:     time.Date(2006, 1, 1, 2, 0, 0, 0, time.UTC),
		TakenSrc:         "exif",
		PhotoTitle:       "",
		TitleSrc:         "",
		PhotoPath:        "2790/02",
		PhotoName:        "Photo01",
		PhotoQuality:     3,
		PhotoResolution:  2,
		PhotoFavorite:    true,
		PhotoPrivate:     false,
		PhotoStory:       false,
		PhotoLat:         48.519234,
		PhotoLng:         9.057997,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraID:         0,
		CameraSerial:     "",
		CameraSrc:        "",
		LensID:           0,
		PlaceID:          "zz",
		LocationID:       "",
		LocationSrc:      "",
		TimeZone:         "",
		PhotoCountry:     "zz",
		PhotoYear:        2790,
		PhotoMonth:       2,
		Description:      Description{},
		DescriptionSrc:   "",
		Camera:           &CameraFixtureEOS6D,
		Lens:             nil,
		Location:         nil,
		Place:            nil,
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{},
		CreatedAt:        time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
	"Photo02": {
		ID:               1000002,
		PhotoUUID:        "pt9jtdre2lvl0yh9",
		TakenAt:          time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC),
		TakenAtLocal:     time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC),
		TakenSrc:         "exif",
		PhotoTitle:       "",
		TitleSrc:         "",
		PhotoPath:        "1990/03",
		PhotoName:        "Photo02",
		PhotoQuality:     3,
		PhotoResolution:  2,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoStory:       false,
		PhotoLat:         48.519234,
		PhotoLng:         9.057997,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraID:         0,
		CameraSerial:     "",
		CameraSrc:        "",
		LensID:           0,
		PlaceID:          "zz",
		LocationID:       "",
		LocationSrc:      "",
		TimeZone:         "",
		PhotoCountry:     "zz",
		PhotoYear:        1990,
		PhotoMonth:       3,
		Description:      Description{},
		DescriptionSrc:   "",
		Camera:           nil,
		Lens:             nil,
		Location:         nil,
		Place:            nil,
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{},
		CreatedAt:        time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
	"Photo03": {
		ID:               1000003,
		PhotoUUID:        "pt9jtdre2lvl0yh0",
		TakenAt:          time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC),
		TakenAtLocal:     time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC),
		TakenSrc:         "exif",
		PhotoTitle:       "",
		TitleSrc:         "",
		PhotoPath:        "1990/04",
		PhotoName:        "Photo02",
		PhotoQuality:     3,
		PhotoResolution:  2,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoStory:       false,
		PhotoLat:         48.519234,
		PhotoLng:         9.057997,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraID:         0,
		CameraSerial:     "",
		CameraSrc:        "",
		LensID:           0,
		PlaceID:          "zz",
		LocationID:       "",
		LocationSrc:      "",
		TimeZone:         "",
		PhotoCountry:     "zz",
		PhotoYear:        1990,
		PhotoMonth:       4,
		Description:      Description{},
		DescriptionSrc:   "",
		Camera:           nil,
		Lens:             nil,
		Location:         nil,
		Place:            nil,
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{},
		CreatedAt:        time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2008, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
	"Photo04": {
		ID:               1000004,
		PhotoUUID:        "pt9jtdre2lvl0y11",
		TakenAt:          time.Date(2014, 7, 17, 15, 42, 12, 0, time.UTC),
		TakenAtLocal:     time.Date(2014, 7, 17, 15, 42, 12, 0, time.UTC),
		TakenSrc:         "exif",
		PhotoTitle:       "Neckarbrücke",
		TitleSrc:         "",
		PhotoPath:        "2014/07",
		PhotoName:        "Photo02",
		PhotoQuality:     3,
		PhotoResolution:  2,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoStory:       false,
		PhotoLat:         48.519234,
		PhotoLng:         9.057997,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraID:         0,
		CameraSerial:     "",
		CameraSrc:        "",
		LensID:           0,
		PlaceID:          "zz",
		LocationID:       "",
		LocationSrc:      "",
		TimeZone:         "",
		PhotoCountry:     "zz",
		PhotoYear:        2014,
		PhotoMonth:       7,
		Description:      Description{},
		DescriptionSrc:   "",
		Camera:           nil,
		Lens:             nil,
		Location:         nil,
		Place:            nil,
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{},
		CreatedAt:        time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
	"Photo05": {
		ID:               1000005,
		PhotoUUID:        "pt9jtdre2lvl0y12",
		TakenAt:          time.Date(2015, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenAtLocal:     time.Date(2015, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenSrc:         "exif",
		PhotoTitle:       "Reunion",
		TitleSrc:         "",
		PhotoPath:        "2014/07",
		PhotoName:        "Photo02",
		PhotoQuality:     3,
		PhotoResolution:  2,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoStory:       false,
		PhotoLat:         -21.342636,
		PhotoLng:         55.466944,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraID:         0,
		CameraSerial:     "",
		CameraSrc:        "",
		LensID:           0,
		PlaceID:          "zz",
		LocationID:       "",
		LocationSrc:      "",
		TimeZone:         "",
		PhotoCountry:     "zz",
		PhotoYear:        2014,
		PhotoMonth:       7,
		Description:      Description{},
		DescriptionSrc:   "",
		Camera:           nil,
		Lens:             nil,
		Location:         nil,
		Place:            nil,
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{},
		CreatedAt:        time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
	"Photo06": {
		ID:               1000006,
		PhotoUUID:        "pt9jtdre2lvl0y13",
		TakenAt:          time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenAtLocal:     time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenSrc:         "exif",
		PhotoTitle:       "ToBeUpdated",
		TitleSrc:         "exif",
		PhotoPath:        "2016/11",
		PhotoName:        "UpdatePhoto",
		PhotoQuality:     0,
		PhotoResolution:  2,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoStory:       false,
		PhotoLat:         -21.342636,
		PhotoLng:         55.466944,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraID:         0,
		CameraSerial:     "",
		CameraSrc:        "",
		LensID:           0,
		PlaceID:          "zz",
		LocationID:       "",
		LocationSrc:      "",
		TimeZone:         "",
		PhotoCountry:     "zz",
		PhotoYear:        2014,
		PhotoMonth:       7,
		Description:      Description{},
		DescriptionSrc:   "",
		Camera:           nil,
		Lens:             nil,
		Location:         nil,
		Place:            nil,
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{},
		CreatedAt:        time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         nil,
		DeletedAt:        nil,
	},
	"Photo07": {
		ID:               1000007,
		PhotoUUID:        "pt9jtdre2lvl0y14",
		TakenAt:          time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenAtLocal:     time.Date(2016, 11, 11, 9, 7, 18, 0, time.UTC),
		TakenSrc:         "",
		PhotoTitle:       "ToBeUpdated",
		TitleSrc:         "exif",
		PhotoPath:        "2016/11",
		PhotoName:        "UpdatePhoto",
		PhotoQuality:     0,
		PhotoResolution:  0,
		PhotoFavorite:    false,
		PhotoPrivate:     false,
		PhotoStory:       false,
		PhotoLat:         -21.342636,
		PhotoLng:         55.466944,
		PhotoAltitude:    0,
		PhotoIso:         0,
		PhotoFocalLength: 0,
		PhotoFNumber:     0,
		PhotoExposure:    "",
		CameraID:         0,
		CameraSerial:     "",
		CameraSrc:        "",
		LensID:           0,
		PlaceID:          "zz",
		LocationID:       "",
		LocationSrc:      "",
		TimeZone:         "",
		PhotoCountry:     "zz",
		PhotoYear:        2014,
		PhotoMonth:       7,
		Description:      Description{},
		DescriptionSrc:   "",
		Camera:           nil,
		Lens:             nil,
		Location:         nil,
		Place:            nil,
		Links:            []Link{},
		Keywords:         []Keyword{},
		Albums:           []Album{},
		Files:            []File{},
		Labels:           []PhotoLabel{},
		CreatedAt:        time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC),
		UpdatedAt:        time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC),
		EditedAt:         &editTime,
		DeletedAt:        nil,
	},
}

var PhotoFixture19800101_000002_D640C559 = PhotoFixtures["19800101_000002_D640C559"]
var PhotoFixturePhoto04 = PhotoFixtures["Photo04"]
var PhotoFixturePhoto01 = PhotoFixtures["Photo01"]
var PhotoFixturePhoto05 = PhotoFixtures["Photo05"]
var PhotoFixturePhoto03 = PhotoFixtures["Photo03"]
var PhotoFixturePhoto06 = PhotoFixtures["Photo06"]
var PhotoFixturePhoto07 = PhotoFixtures["Photo07"]

// CreatePhotoFixtures inserts known entities into the database for testing.
func CreatePhotoFixtures() {
	for _, entity := range PhotoFixtures {
		Db().Create(&entity)
	}
}
