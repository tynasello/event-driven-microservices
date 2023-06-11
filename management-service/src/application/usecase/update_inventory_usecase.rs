use std::collections::HashMap;

use std::env;

use crate::application::interfaces::i_rest_service::IRestService;

pub struct UpdateInventoryUsecase<'a> {
    rest_service: &'a Box<dyn IRestService>,
}

impl<'a> UpdateInventoryUsecase<'a> {
    pub fn new(rest_service: &'a Box<dyn IRestService>) -> Self {
        Self { rest_service }
    }

    pub async fn execute(&self, label: &str, add_to_quantity: i32) -> Result<String, String> {
        let inventory_service_uri: &str =
            &(env::var("INVENTORY_SERVICE_URI").unwrap() + "/update-inventory-item");

        let add_to_quantity_string: &str = &add_to_quantity.to_string();

        let mut map = HashMap::new();
        map.insert("label", label);
        map.insert("addToQuantity", add_to_quantity_string);

        let response_result = self
            .rest_service
            .fetch("POST", inventory_service_uri, "", map)
            .await;

        return match response_result {
            Ok(response) => Ok(response.body),
            Err(error) => Err(error),
        };
    }
}
