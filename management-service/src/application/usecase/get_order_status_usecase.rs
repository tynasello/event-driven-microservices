use std::collections::HashMap;

use std::env;

use crate::application::interfaces::i_rest_service::IRestService;

pub struct GetOrderStatusUsecase<'a> {
    rest_service: &'a dyn IRestService,
}

impl<'a> GetOrderStatusUsecase<'a> {
    pub fn new(rest_service: &'a dyn IRestService) -> Self {
        Self { rest_service }
    }

    pub async fn execute(&self, access_token: String, id: i32) -> Result<String, String> {
        let inventory_service_uri: &str = &(env::var("ORDER_SERVICE_URI").unwrap() + "/get-order");

        let id_string: &str = &id.to_string();

        let mut map = HashMap::new();
        map.insert("id", id_string);

        let response_result = self
            .rest_service
            .fetch("GET", inventory_service_uri, &access_token, map)
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
