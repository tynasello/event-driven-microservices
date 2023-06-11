use clap::ArgMatches;

use crate::application::{
    interfaces::i_rest_service::IRestService, usecase::login_usecase::LoginUsecase,
};

pub async fn login_subcommand<'a>(
    login_matches: &'a ArgMatches<'_>,
    rest_service: &'a Box<dyn IRestService>,
) {
    let username = login_matches.value_of("username").unwrap();
    let password = login_matches.value_of("password").unwrap();

    let login_user_service = LoginUsecase::new(&rest_service);
    let login_user_result = login_user_service.execute(username, password).await;
    if let Ok(access_token) = login_user_result {
        println!("Access token: {}", access_token);
    } else {
        println!("Error logging in: {}", login_user_result.unwrap_err());
    }
}
