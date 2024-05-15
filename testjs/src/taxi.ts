export const Taxi = () => 55;

export function CalculateFare(km: number, minute: number): number {
  const minimumFare = 35;
  const roundedKm = RoundKm(km);
  const roundedMinute = Math.ceil(minute);

  const result = 4 * roundedKm + 1 * roundedMinute;

  return result >= minimumFare ? result : minimumFare;
}

export function RoundKm(input: number): number {
  return Math.ceil(input * 2) / 2;
}
