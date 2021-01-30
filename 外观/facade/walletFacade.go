package facade

import "fmt"

type WalletFacade struct {
    account      *Account
    wallet       *Wallet
    securityCode *SecurityCode
    notification *Notification
    ledger       *Ledger
}

func NewWalletFacade(accountID string, code int) *WalletFacade {
    fmt.Println("Starting create account")
    walletFacacde := &WalletFacade{
        account:      NewAccount(accountID),
        securityCode: NewSecurityCode(code),
        wallet:       NewWallet(),
        notification: &Notification{},
        ledger:       &Ledger{},
    }
    fmt.Println("Account created")
    return walletFacacde
}

func (w *WalletFacade) AddMoneyToWallet(accountID string, securityCode int, amount int) error {
    fmt.Println("Starting add money to wallet")
    err := w.account.CheckAccount(accountID)
    if err != nil {
        return err
    }
    err = w.securityCode.CheckCode(securityCode)
    if err != nil {
        return err
    }
    w.wallet.CreditBalance(amount)
    w.notification.SendWalletCreditNotification()
    w.ledger.MakeEntry(accountID, "credit", amount)
    return nil
}

func (w *WalletFacade) DeductMoneyFromWallet(accountID string, securityCode int, amount int) error {
    fmt.Println("Starting debit money from wallet")
    err := w.account.CheckAccount(accountID)
    if err != nil {
        return err
    }

    err = w.securityCode.CheckCode(securityCode)
    if err != nil {
        return err
    }
    err = w.wallet.DebitBalance(amount)
    if err != nil {
        return err
    }
    w.notification.SendWalletDebitNotification()
    w.ledger.MakeEntry(accountID, "credit", amount)
    return nil
}