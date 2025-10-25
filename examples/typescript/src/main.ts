import { createGrpcTransport } from "@connectrpc/connect-node";
import { createClient } from "@connectrpc/connect";
import type {} from "@bufbuild/protobuf";
import { SnowflakeService } from "@snowflake/proto/snowflake/v1/snowflake_service_pb";

const transport = createGrpcTransport({
  baseUrl: "http://localhost:8081",
});

const client = createClient(SnowflakeService, transport);

(async () => {
    const response = await client.nextSnowflake({});

    console.log("Created a Snowflake:", response.snowflake?.stringValue);
})();
