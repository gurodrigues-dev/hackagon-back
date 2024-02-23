package types

import "fmt"

const (
	TITLE_LIMIT = 100
	DESC_LIMIT  = 300
)

func (q *QuestionCreateRequest) ValidateCreate() error {

	if len(q.Title) == 0 {
		return fmt.Errorf("question title cannot be empty")
	}

	if len(q.Title) > TITLE_LIMIT {
		return fmt.Errorf("question title too big. Limit is %d", TITLE_LIMIT)
	}

	if len(q.Description) == 0 {
		return fmt.Errorf("question description cannot be empty")
	}

	if len(q.Description) > DESC_LIMIT {
		return fmt.Errorf("question description too big. Limit is %d", DESC_LIMIT)
	}

	return nil

}
