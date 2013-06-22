package campfire

type Account struct {
    conn *Connection

    Id          int    `json:"id"`
    Name        string `json:"name"`
    Subdomain   string `json:"subdomain"`
    OwnerId     int    `json:"owner_id"`
    Plan        string `json:"plan"`
    StorageUsed int    `json:"storage"`
    Timezone    string `json:"time_zone"`
    CreatedAt   string `json:"created_at"`
    UpdatedAt   string `json:"updated_at"`
}

type AccountResult struct {
    Account *Account `json:"account"`
}
