use std::collections::HashMap;

use async_trait::async_trait;
use reqwest::{
    header::{HeaderMap, HeaderValue, COOKIE},
    Client,
};

use crate::application::interfaces::i_rest_service::{IRestService, RestServiceResponse};

// http_client should be some interface
pub struct RestService {
    http_client: reqwest::Client,
}

impl RestService {
    pub fn new(http_client: Option<Client>) -> Self {
        let http_client = http_client.unwrap_or_else(Client::new);
        Self { http_client }
    }
}

#[async_trait]
impl IRestService for RestService {
    async fn fetch(
        &self,
        method: &str,
        endpoint: &str,
        access_token: &str,
        body: HashMap<&str, &str>,
    ) -> Result<RestServiceResponse, String> {
        let mut request_headers = HeaderMap::new();
        let access_token_cookie: &str = &["access-token=", access_token].concat();
        request_headers.insert(COOKIE, HeaderValue::from_str(access_token_cookie).unwrap());

        let response = if method == "POST" {
            self.http_client
                .post(endpoint)
                .headers(request_headers)
                .json(&body)
                .send()
                .await
        } else if method == "GET" {
            self.http_client
                .get(endpoint)
                .headers(request_headers)
                .json(&body)
                .send()
                .await
        } else {
            return Err("Invalid method for http request".to_string());
        };

        match response {
            Ok(response) => {
                if response.status().is_success() == false {
                    let response_body = response.text().await.unwrap_or_else(|_| "".to_string());
                    return Err(response_body);
                }

                let response_headers = response.headers();
                let raw_cookies = response_headers
                    .get("set-cookie")
                    .map(|c| c.to_str().unwrap_or(""))
                    .unwrap_or("");

                let access_token = raw_cookies
                    .split(';')
                    .find(|s| s.starts_with("access-token="))
                    .map(|s| s.split('=').skip(1).next().unwrap_or(""))
                    .map(|s| s.trim())
                    .unwrap_or("");

                let mut response_cookies = HashMap::new();
                response_cookies.insert("access-token".to_string(), access_token.to_string());

                let response_body = response.text().await.unwrap_or_else(|_| "".to_string());

                let rest_service_response = RestServiceResponse {
                    body: response_body,
                    cookies: response_cookies,
                };

                return Ok(rest_service_response);
            }
            Err(error) => {
                return Err(error.to_string());
            }
        }
    }
}
