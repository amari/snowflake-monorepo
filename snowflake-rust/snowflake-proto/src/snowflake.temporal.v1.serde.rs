// @generated
impl serde::Serialize for BatchNextSnowflakeActivityInput {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.count != 0 {
            len += 1;
        }
        if self.wait {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("snowflake.temporal.v1.BatchNextSnowflakeActivityInput", len)?;
        if self.count != 0 {
            struct_ser.serialize_field("count", &self.count)?;
        }
        if self.wait {
            struct_ser.serialize_field("wait", &self.wait)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for BatchNextSnowflakeActivityInput {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "count",
            "wait",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Count,
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
                            "count" => Ok(GeneratedField::Count),
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
            type Value = BatchNextSnowflakeActivityInput;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct snowflake.temporal.v1.BatchNextSnowflakeActivityInput")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<BatchNextSnowflakeActivityInput, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut count__ = None;
                let mut wait__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Count => {
                            if count__.is_some() {
                                return Err(serde::de::Error::duplicate_field("count"));
                            }
                            count__ = 
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
                Ok(BatchNextSnowflakeActivityInput {
                    count: count__.unwrap_or_default(),
                    wait: wait__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("snowflake.temporal.v1.BatchNextSnowflakeActivityInput", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for BatchNextSnowflakeActivityOutput {
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
        let mut struct_ser = serializer.serialize_struct("snowflake.temporal.v1.BatchNextSnowflakeActivityOutput", len)?;
        if !self.snowflakes.is_empty() {
            struct_ser.serialize_field("snowflakes", &self.snowflakes)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for BatchNextSnowflakeActivityOutput {
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
            type Value = BatchNextSnowflakeActivityOutput;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct snowflake.temporal.v1.BatchNextSnowflakeActivityOutput")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<BatchNextSnowflakeActivityOutput, V::Error>
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
                Ok(BatchNextSnowflakeActivityOutput {
                    snowflakes: snowflakes__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("snowflake.temporal.v1.BatchNextSnowflakeActivityOutput", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for NextSnowflakeActivityInput {
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
        let mut struct_ser = serializer.serialize_struct("snowflake.temporal.v1.NextSnowflakeActivityInput", len)?;
        if self.wait {
            struct_ser.serialize_field("wait", &self.wait)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for NextSnowflakeActivityInput {
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
            type Value = NextSnowflakeActivityInput;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct snowflake.temporal.v1.NextSnowflakeActivityInput")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<NextSnowflakeActivityInput, V::Error>
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
                Ok(NextSnowflakeActivityInput {
                    wait: wait__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("snowflake.temporal.v1.NextSnowflakeActivityInput", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for NextSnowflakeActivityOutput {
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
        let mut struct_ser = serializer.serialize_struct("snowflake.temporal.v1.NextSnowflakeActivityOutput", len)?;
        if let Some(v) = self.snowflake.as_ref() {
            struct_ser.serialize_field("snowflake", v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for NextSnowflakeActivityOutput {
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
            type Value = NextSnowflakeActivityOutput;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct snowflake.temporal.v1.NextSnowflakeActivityOutput")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<NextSnowflakeActivityOutput, V::Error>
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
                Ok(NextSnowflakeActivityOutput {
                    snowflake: snowflake__,
                })
            }
        }
        deserializer.deserialize_struct("snowflake.temporal.v1.NextSnowflakeActivityOutput", FIELDS, GeneratedVisitor)
    }
}
