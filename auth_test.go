package resharmonics

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/semaphore"
)

func TestAuthLockSimple(t *testing.T) {

	// arrange
	mockSemaphore := semaphore.NewWeighted(1)

	// act
	hasToDo, errLocking := hasToDoAuth(mockSemaphore)

	couldRelock := mockSemaphore.TryAcquire(1)

	// verify
	assert.Nil(t, errLocking)
	assert.True(t, hasToDo)
	assert.False(t, couldRelock)
}

func TestSecondCannotRelock(t *testing.T) {

	// arrange
	mockSemaphore := semaphore.NewWeighted(1)

	// act
	hasToDo1, errLocking1 := hasToDoAuth(mockSemaphore)
	couldRelock1 := mockSemaphore.TryAcquire(1)

	secondStatementIssued := false

	var hasToDo2 bool
	var errLocking2 error

	go func() {
		hasToDo2, errLocking2 = hasToDoAuth(mockSemaphore)
		secondStatementIssued = true
	}()

	time.Sleep(time.Millisecond * 400)
	couldRelock2 := mockSemaphore.TryAcquire(1)
	mockSemaphore.Release(1)

	for {
		if secondStatementIssued == false {
			time.Sleep(time.Millisecond * 100)
			continue
		} else {
			break
		}
	}

	// verify
	assert.Nil(t, errLocking1)
	assert.True(t, hasToDo1)
	assert.False(t, couldRelock1)
	assert.False(t, couldRelock2)

	assert.Nil(t, errLocking2)
	assert.False(t, hasToDo2)

}
