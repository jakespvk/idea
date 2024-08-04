const std = @import("std");
const Property = struct {
    const address: [100]u8 = undefined;
};

const property1: Property = "145 Belle Meade Pl";
const property2: Property = "1860 Aquarius St";
const property3: Property = "177 Coach Light Ln";

const db: std.ArrayList(Property).init(std.testing.allocator);
defer db.deinit();

try db.append(property1);
try db.append(property2);
try db.append(property3);

pub fn search(input: [100]u8) Property {
    // binary search? search for property in db
    // returns property *item*
    for (db) |record| {
        if (record == input) {
            return record;
        }
    }
    return Property.None;
}
