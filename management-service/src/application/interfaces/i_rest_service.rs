use std::collections::HashMap;

use async_trait::async_trait;

pub struct RestServiceResponse {
    pub body: String,
    pub cookies: HashMap<String, String>,
}

#[async_trait]
pub trait IRestService {
    async fn fetch(
        &self,
        endpoint: &str,
        access_token: &str,
        body: HashMap<&str, &str>,
    ) -> Result<RestServiceResponse, String>;
}
