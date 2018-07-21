package components

import "testing"

func TestGetUID(t *testing.T) {
	firstUID := getUID()

	count := 10
	for count > 0 {
		count--
		currentUID := getUID()
		if currentUID != firstUID {
			t.Errorf("%d current get uid %s != fistUID %s", count, currentUID, firstUID)
			return
		}
	}
	t.Logf("get same uid in different times, uid is %s", firstUID)
}
