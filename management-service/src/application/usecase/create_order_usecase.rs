use std::collections::HashMap;

use std::env;

use crate::application::interfaces::i_rest_service::IRestService;

pub struct CreateOrderUsecase<'a> {
    rest_service: &'a Box<dyn IRestService>,
}

impl<'a> CreateOrderUsecase<'a> {
    pub fn new(rest_service: &'a Box<dyn IRestService>) -> Self {
        Self { rest_service }
    }

    pub async fn execute(
        &self,
        access_token: String,
        product_name: &str,
        product_quantity: i32,
    ) -> Result<String, String> {
        let inventory_service_uri: &str =
            &(env::var("INVENTORY_SERVICE_URI").unwrap() + "/create-order");

        let product_quantity_string: &str = &product_quantity.to_string();

        let mut map = HashMap::new();
        map.insert("productName", product_name);
        map.insert("produtQuantity", product_quantity_string);

        let response_result = self
            .rest_service
            .fetch("POST", inventory_service_uri, &access_token, map)
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
