import asyncio

from grpclib.client import Channel

import snowflake_python.proto.snowflake.v1 as snowflake_v1

async def main():
    channel = Channel(host="localhost", port=8081)
    try:
        stub = snowflake_v1.SnowflakeServiceStub(channel)
        response = await stub.next_snowflake(wait=False)
        print("Created a Snowflake:", response.snowflake.string_value)
    finally:
        channel.close()

if __name__ == "__main__":
    loop = asyncio.new_event_loop()
    asyncio.set_event_loop(loop)
    try:
        loop.run_until_complete(main())
    except KeyboardInterrupt:
        pass
    finally:
        loop.close()
