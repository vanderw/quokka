local m = require("mymodule")

-- calling go from lua
print(m.name, m.myfunc())

print('Hello')
print("i'm going to call double()", double(2, 5))

-- l_double will be called from golang
function l_double(v)
    return 2 * v
end

-- userdata
local p = person.new('john')
print('name:' .. p:name())
-- change name
p:name('lucy')
print('new name:' .. p:name())
