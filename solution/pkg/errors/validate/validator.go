package validate

import "context"

type Condition func(ctx context.Context) error

func Validate(ctx context.Context, conditions ...Condition) error {
	ve := NewValidationError()

	for _, condition := range conditions {
		err := condition(ctx)
		if err != nil {
			if IsValidationError(err) {
				ve.Add(err.Error())
				continue
			}

			return err
		}
	}

	if len(ve.messages) == 0 {
		return nil
	}

	return ve
}
