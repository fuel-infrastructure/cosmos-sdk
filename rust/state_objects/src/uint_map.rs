use interchain_core::{Context, Response};
use crate::codec::ObjectKey;
use crate::Map;

/// A map from keys to 128-bit unsigned integers.
pub struct UInt128Map<K> {
    map: Map<K, u128>,
}

impl<'a, K: ObjectKey<'a>> UInt128Map<K> {
    /// Gets the current value for the given key, defaulting always to 0.
    pub fn get(&self, ctx: &Context, key: K::Value) -> Response<u128> {
        // let value = self.map.get(ctx, key)?;
        // Ok(value.unwrap_or(0))
        todo!()
    }

    /// Adds the given value to the current value for the given key.
    pub fn add(&self, ctx: &mut Context, key: K::Value, value: u128) -> Response<u128> {
        todo!()
    }

    /// Subtracts the given value from the current value for the given key,
    /// returning an error if the subtraction would result in a negative value.
    pub fn safe_sub(&self, ctx: &mut Context, key: K::Value, value: u128) -> Response<u128> {
        todo!()
    }
}