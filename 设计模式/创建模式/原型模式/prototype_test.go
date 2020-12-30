package prototype

import "testing"

func TestPrototype(t *testing.T) {
	resume := NewResume()

	resume.setPersonalInfo("路飞", "男", "18")
	resume.setWorkExperience("3", "google")

	cloneResume := resume.clone()
	cloneResume.setPersonalInfo("索隆", "男", "21")
	cloneResume.setWorkExperience("2", "Apple")
	resume.display()
	cloneResume.display()
}
