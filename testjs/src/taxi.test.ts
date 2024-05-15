import { CalculateFare, RoundKm, Taxi } from "./taxi";

describe("Taxi", () => {
  test("Should return 0 when km and minute are 0", () => {
    const km = 0;
    const minute = 0;

    const result = CalculateFare(km, minute);

    expect(result).toEqual(35);
  });

  test("Should return 5 when km and minute are 1", () => {
    const km = 1;
    const minute = 1;

    const result = CalculateFare(km, minute);

    expect(result).toEqual(35);
  });

  test("Should return 52 when km is 10.1 and minute is 10", () => {
    const km = 10.1;
    const minute = 10;

    const result = CalculateFare(km, minute);

    expect(result).toEqual(52);
  });

  test("Should return 48 when km is 10.1 and minute is 5.5", () => {
    const km = 10.1;
    const minute = 5.5;

    const result = CalculateFare(km, minute);

    expect(result).toEqual(48);
  });

  test("Should return 50 when km is 10.7 and minute is 5.5", () => {
    const km = 10.7;
    const minute = 5.5;

    const result = CalculateFare(km, minute);

    expect(result).toEqual(50);
  });

  test("Should return minumumFare (35) when result is lower than minimumFare", () => {
    const km = 1;
    const minute = 1;

    const result = CalculateFare(km, minute);

    expect(result).toEqual(35);
  });
});

describe("Test RoundKM function", () => {
  test("Should return 10 rounded km correctly when km = 10", () => {
    const km = 10;

    const result = RoundKm(km);

    expect(result).toEqual(10);
  });

  test("Should return 10.5 rounded km correctly when km = 10.1", () => {
    const km = 10.1;

    const result = RoundKm(km);

    expect(result).toEqual(10.5);
  });
});

// 4 * km + 1 * minute
