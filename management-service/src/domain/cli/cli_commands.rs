pub fn get_cli_commands() -> Vec<(&'static str, Vec<(&'static str, &'static str)>)> {
    vec![
        (
            "login",
            vec![
                ("username", "Username to login with"),
                ("password", "Password for username"),
            ],
        ),
        (
            "signup",
            vec![
                ("username", "Username to sign up with"),
                ("password", "Password to sign up with"),
            ],
        ),
        (
            "create-order",
            vec![
                ("access-token", "Use to authenticate the user"),
                ("product-name", "Sets the product name for order"),
                ("product-quantity", "Sets the product quantity for order"),
            ],
        ),
        (
            "get-order-status",
            vec![
                ("access-token", "Use to authenticate the user"),
                ("order-id", "Sets the order id to get status for"),
            ],
        ),
        (
            "add-inventory",
            vec![
                ("label", "Sets the product name for inventory"),
                ("quantity", "Sets the product quantity for inventory"),
            ],
        ),
        (
            "update-inventory",
            vec![
                ("label", "Sets the product name for inventory"),
                ("quantity-to-add", "Sets the product quantity for inventory"),
            ],
        ),
    ]
}
