use std::collections::HashMap;

use std::env;

use crate::application::interfaces::i_rest_service::IRestService;

pub struct LoginUsecase<'a> {
    rest_service: &'a Box<dyn IRestService>,
}

impl<'a> LoginUsecase<'a> {
    pub fn new(rest_service: &'a Box<dyn IRestService>) -> Self {
        Self { rest_service }
    }

    pub async fn execute(&self, username: &str, password: &str) -> Result<String, String> {
        let user_service_uri: &str = &(env::var("USER_SERVICE_URI").unwrap() + "/login");

        let mut map = HashMap::new();
        map.insert("username", username);
        map.insert("password", password);

        let response_result = self
            .rest_service
            .fetch("GET", user_service_uri, "", map)
            .await;

        match response_result {
            Ok(response) => response
                .cookies
                .get("access_token")
                .map(|cookie| cookie.to_string())
                .ok_or("".to_string()),
            Err(error) => Err(error),
        }
    }
}
