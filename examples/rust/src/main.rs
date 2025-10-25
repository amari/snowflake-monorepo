use snowflake_proto;
use snowflake_proto::snowflake::v1::NextSnowflakeRequest;
use snowflake_proto::snowflake::v1::snowflake_service_client::SnowflakeServiceClient;
use tokio;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let mut client = SnowflakeServiceClient::connect("http://[::1]:8081").await?;

    let resp = client
        .next_snowflake(NextSnowflakeRequest { wait: false })
        .await?;

    println!(
        "Created a Snowflake: {}",
        &resp.into_inner().snowflake.unwrap().string_value
    );

    Ok(())
}
