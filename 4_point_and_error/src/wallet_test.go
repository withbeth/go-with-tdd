package main

import "testing"

func TestWallet(t *testing.T) {

	assertBitCoin := func(t *testing.T, expected BitCoin, actual BitCoin) {
		t.Helper()
		if expected != actual {
			t.Errorf("Expected : %s, but Actual : %s", expected, actual)
		}
	}

	t.Run("Should can access Wallet balance directly within same package", func(t *testing.T) {

		wallet := Wallet{10}

		expected := wallet.balance
		actual := BitCoin(10)

		assertBitCoin(t, expected, actual)
	})

	t.Run("Should match with given Deposit and Balance by using constructor", func(t *testing.T) {
		expected := BitCoin(10)
		wallet := Wallet{expected}
		actual := wallet.Balance()

		assertBitCoin(t, expected, actual)
	})

	t.Run("Should match with given Deposit and Balance by calling SET method", func(t *testing.T) {
		wallet := Wallet{}

		expected := BitCoin(10)
		wallet.Deposit(expected)
		actual := wallet.Balance()

		assertBitCoin(t, expected, actual)
	})

	t.Run("Should match after withdraw", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(BitCoin(15))
		actual , err := wallet.Withdraw(BitCoin(5))

		expected := BitCoin(10)
		assertBitCoin(t, expected, actual)
		if err != nil {
			t.Errorf("Expected : nil, but Actual : %v", err)
		}

	})


	withDrawTestCases := [] struct{
		name string
		deposit BitCoin
		withdraw BitCoin
		doesShouldHaveError bool
		expectedBalance BitCoin
	}{
		{ "Should have an error and keep origin balance when withdraw more than balance", BitCoin(5), BitCoin(10), true, BitCoin(5)},
		{ "Should not have an error and take out balance when withdraw just all balance", BitCoin(10), BitCoin(10), false, BitCoin(0)},
		{ "Should not have an error and take out balance when withdraw less than balance", BitCoin(10), BitCoin(9), false, BitCoin(1)},
	}

	for _, testCase := range withDrawTestCases {
		t.Run(testCase.name, func(t *testing.T) {
			wallet := Wallet{}
			wallet.Deposit(testCase.deposit)

			actual , err := wallet.Withdraw(testCase.withdraw)

			if testCase.expectedBalance != actual {
				assertBitCoin(t, testCase.expectedBalance, actual)
			}

			if testCase.doesShouldHaveError && err == nil {
				t.Errorf("Expected : OverDraftError, but Actual : nil")
			} else if !testCase.doesShouldHaveError && err != nil {
				t.Errorf("Expected : nil, but Actual : %v", err)
			}
			if err != nil {
				t.Logf("Got Error: %s", err.Error())
			}
		})
	}
}
