export function fetcherJson(url: string) {
  return fetch(url).then((res) => res.json());
}
