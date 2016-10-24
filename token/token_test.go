package token

import "testing"

func TestPrecedence(t *testing.T) {
	for tok := token_beg + 1; tok < token_end; tok++ {
		if !tok.IsOperator() && tok.Precedence() > 0 {
			t.Errorf("%s should have the lowest precedence.", tok.String())
		}
	}

	if OR.Precedence() <= 0 {
		t.Errorf("OR should have higher than lowest precedence.")
	}

	if AND.Precedence() <= OR.Precedence() {
		t.Errorf("AND should have higher precedence than OR.")
	}

	if EQL.Precedence() != NEQ.Precedence() ||
		EQL.Precedence() != LSS.Precedence() ||
		EQL.Precedence() != LEQ.Precedence() ||
		EQL.Precedence() != GTR.Precedence() ||
		EQL.Precedence() != GEQ.Precedence() {
		t.Errorf("EQL, NEQ, LSS, LEQ, GTR, GEQ should have the same precedence.")
	}
	if EQL.Precedence() <= AND.Precedence() {
		t.Errorf("EQL should have higher precedence than AND.")
	}

	if ADD.Precedence() != SUB.Precedence() {
		t.Errorf("ADD and SUB should have the same precedence.")
	}
	if ADD.Precedence() <= EQL.Precedence() {
		t.Errorf("ADD should have higher precedence than EQL.")
	}

	if MUL.Precedence() != QUO.Precedence() ||
		MUL.Precedence() != REM.Precedence() {
		t.Errorf("MUL, QUO, REM should have the same precedence.")
	}
	if MUL.Precedence() <= ADD.Precedence() {
		t.Errorf("MUL should have higher precedence than ADD.")
	}
}

func TestIsLiteral(t *testing.T) {
	for tok := token_beg + 1; tok < token_end; tok++ {
		if (tok <= literal_beg || tok >= literal_end) && tok.IsLiteral() {
			t.Errorf("%s should not be a literal.", tok.String())
		}
		if (tok > literal_beg && tok < literal_end) && !tok.IsLiteral() {
			t.Errorf("%s should be a literal.", tok.String())
		}
	}
}

func TestIsOperator(t *testing.T) {
	for tok := token_beg + 1; tok < token_end; tok++ {
		if (tok <= operator_beg || tok >= operator_end) && tok.IsOperator() {
			t.Errorf("%s should not be an operator.", tok.String())
		}
		if (tok > operator_beg && tok < operator_end) && !tok.IsOperator() {
			t.Errorf("%s should be an operator.", tok.String())
		}
	}
}
