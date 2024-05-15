export const Taxi = () => 55;

export function CalculateFare(km: number, minute: number): number {
  const roundedKm = Math.ceil(km * 2) / 2;
  const roundedMinute = Math.ceil(minute);

  const result = 4 * roundedKm + 1 * roundedMinute;

  return result;
}
