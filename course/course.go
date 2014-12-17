package courses

import (
	"moodle"
	"io"
)

type Course struct {
	name string
	id int
}

func parseCourse(body io.ReadCloser) *Course {

	return nil
}

func GetCourses(client *moodle.Client) {

}

func GetCourseFromId(client *moodle.Client, id int) {

}

func GetCourseFromName(client *moodle.Client, name string) {

}