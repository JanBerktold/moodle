package user

import (
	"moodle"
	"course"
)

type User struct {
	name, email, country, city string
	id int
	courses []*Course
}

// Gets a 
func (user *User) UpdateCourses(client *moodle.Client) {

}


// GetCourses() returns a slice of the courses of which the user is a member of.
// If there is a cached slice available, then it is returned otherwise UpdateCourses() and its result is returned.
func (user *User) GetCourses() []*Course {

	if (user.courses == nil) {
		user.UpdateCourses()
	}

	return user.courses
}