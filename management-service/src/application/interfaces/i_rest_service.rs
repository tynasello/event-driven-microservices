use std::collections::HashMap;

use async_trait::async_trait;
use mockall::automock;

#[derive(Debug)]
pub struct RestServiceResponse {
    pub body: String,
    pub cookies: HashMap<String, String>,
}

#[async_trait]
#[automock]
pub trait IRestService {
    async fn fetch(
        &self,
        method: &str,
        endpoint: &str,
        access_token: &str,
        body: HashMap<&str, &str>,
    ) -> Result<RestServiceResponse, String>;
}
