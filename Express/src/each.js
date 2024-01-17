export default function each(items, fn) {
  return items.map((item) => fn(item)).join("");
}
