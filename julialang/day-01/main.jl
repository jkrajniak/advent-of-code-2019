function get_fuel(mass)
    return floor(mass / 3.0) - 2
end

function get_fuel_extended(mass)
    fuel = get_fuel(mass)
    total_fuel = fuel
    while true
        fuel = get_fuel(fuel)
        if fuel <= 0
            break
        end
        total_fuel += fuel
    end
    return total_fuel
end


fname = "input.txt"
total_fuel = 0
open(fname, "r") do f
    for line in eachline(f)
        mass = parse(Int, line)
        global total_fuel += get_fuel_extended(mass)
    end
end
println(Int(total_fuel))