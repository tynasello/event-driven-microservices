use std::collections::HashMap;

use std::env;

use crate::application::interfaces::i_rest_service::IRestService;

pub struct AddInventoryUsecase<'a> {
    rest_service: &'a dyn IRestService,
}

impl<'a> AddInventoryUsecase<'a> {
    pub fn new(rest_service: &'a dyn IRestService) -> Self {
        Self { rest_service }
    }

    pub async fn execute(&self, label: &str, quantity_in_stock: i32) -> Result<String, String> {
        let inventory_service_uri: &str =
            &(env::var("INVENTORY_SERVICE_URI").unwrap() + "/add-inventory-item");

        let quantity_in_stock_string: &str = &quantity_in_stock.to_string();

        let mut map = HashMap::new();
        map.insert("label", label);
        map.insert("quantityInStock", quantity_in_stock_string);
        map.insert("quantityReserved", "0");

        let response_result = self
            .rest_service
            .fetch("POST", inventory_service_uri, "", map)
            .await;

        match response_result {
            Ok(response) => {
                return Ok(response.body);
            }
            Err(error) => {
                return Err(error);
            }
        }
    }
}
