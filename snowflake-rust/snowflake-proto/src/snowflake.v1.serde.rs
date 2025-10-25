// @generated
impl serde::Serialize for BatchNextSnowflakeRequest {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.batch_size != 0 {
            len += 1;
        }
        if self.wait {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("snowflake.v1.BatchNextSnowflakeRequest", len)?;
        if self.batch_size != 0 {
            struct_ser.serialize_field("batchSize", &self.batch_size)?;
        }
        if self.wait {
            struct_ser.serialize_field("wait", &self.wait)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for BatchNextSnowflakeRequest {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "batch_size",
            "batchSize",
            "wait",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            BatchSize,
            Wait,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "batchSize" | "batch_size" => Ok(GeneratedField::BatchSize),
                            "wait" => Ok(GeneratedField::Wait),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = BatchNextSnowflakeRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct snowflake.v1.BatchNextSnowflakeRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<BatchNextSnowflakeRequest, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut batch_size__ = None;
                let mut wait__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::BatchSize => {
                            if batch_size__.is_some() {
                                return Err(serde::de::Error::duplicate_field("batchSize"));
                            }
                            batch_size__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Wait => {
                            if wait__.is_some() {
                                return Err(serde::de::Error::duplicate_field("wait"));
                            }
                            wait__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(BatchNextSnowflakeRequest {
                    batch_size: batch_size__.unwrap_or_default(),
                    wait: wait__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("snowflake.v1.BatchNextSnowflakeRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for BatchNextSnowflakeResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.snowflakes.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("snowflake.v1.BatchNextSnowflakeResponse", len)?;
        if !self.snowflakes.is_empty() {
            struct_ser.serialize_field("snowflakes", &self.snowflakes)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for BatchNextSnowflakeResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "snowflakes",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Snowflakes,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "snowflakes" => Ok(GeneratedField::Snowflakes),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = BatchNextSnowflakeResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct snowflake.v1.BatchNextSnowflakeResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<BatchNextSnowflakeResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut snowflakes__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Snowflakes => {
                            if snowflakes__.is_some() {
                                return Err(serde::de::Error::duplicate_field("snowflakes"));
                            }
                            snowflakes__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(BatchNextSnowflakeResponse {
                    snowflakes: snowflakes__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("snowflake.v1.BatchNextSnowflakeResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for NextSnowflakeRequest {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.wait {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("snowflake.v1.NextSnowflakeRequest", len)?;
        if self.wait {
            struct_ser.serialize_field("wait", &self.wait)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for NextSnowflakeRequest {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "wait",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Wait,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "wait" => Ok(GeneratedField::Wait),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = NextSnowflakeRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct snowflake.v1.NextSnowflakeRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<NextSnowflakeRequest, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut wait__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Wait => {
                            if wait__.is_some() {
                                return Err(serde::de::Error::duplicate_field("wait"));
                            }
                            wait__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(NextSnowflakeRequest {
                    wait: wait__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("snowflake.v1.NextSnowflakeRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for NextSnowflakeResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.snowflake.is_some() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("snowflake.v1.NextSnowflakeResponse", len)?;
        if let Some(v) = self.snowflake.as_ref() {
            struct_ser.serialize_field("snowflake", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for NextSnowflakeResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "snowflake",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Snowflake,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "snowflake" => Ok(GeneratedField::Snowflake),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = NextSnowflakeResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct snowflake.v1.NextSnowflakeResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<NextSnowflakeResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut snowflake__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Snowflake => {
                            if snowflake__.is_some() {
                                return Err(serde::de::Error::duplicate_field("snowflake"));
                            }
                            snowflake__ = map_.next_value()?;
                        }
                    }
                }
                Ok(NextSnowflakeResponse {
                    snowflake: snowflake__,
                })
            }
        }
        deserializer.deserialize_struct("snowflake.v1.NextSnowflakeResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for Snowflake {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.int64_value != 0 {
            len += 1;
        }
        if !self.string_value.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("snowflake.v1.Snowflake", len)?;
        if self.int64_value != 0 {
            #[allow(clippy::needless_borrow)]
            #[allow(clippy::needless_borrows_for_generic_args)]
            struct_ser.serialize_field("int64Value", ToString::to_string(&self.int64_value).as_str())?;
        }
        if !self.string_value.is_empty() {
            struct_ser.serialize_field("stringValue", &self.string_value)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for Snowflake {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "int64_value",
            "int64Value",
            "string_value",
            "stringValue",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Int64Value,
            StringValue,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "int64Value" | "int64_value" => Ok(GeneratedField::Int64Value),
                            "stringValue" | "string_value" => Ok(GeneratedField::StringValue),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = Snowflake;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct snowflake.v1.Snowflake")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<Snowflake, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut int64_value__ = None;
                let mut string_value__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Int64Value => {
                            if int64_value__.is_some() {
                                return Err(serde::de::Error::duplicate_field("int64Value"));
                            }
                            int64_value__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::StringValue => {
                            if string_value__.is_some() {
                                return Err(serde::de::Error::duplicate_field("stringValue"));
                            }
                            string_value__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(Snowflake {
                    int64_value: int64_value__.unwrap_or_default(),
                    string_value: string_value__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("snowflake.v1.Snowflake", FIELDS, GeneratedVisitor)
    }
}
