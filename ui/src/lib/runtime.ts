// copied from https://stackoverflow.com/questions/5639346/what-is-the-shortest-function-for-reading-a-cookie-by-name-in-javascript
const getCookieValue = (name: string) => {
  const entry = document.cookie.match('(^|[^;]+)\\s*' + name + '\\s*=\\s*([^;]+)');
  return entry ? entry.pop() : undefined;
};

const getLastName = (): string => {
  const lastName = getCookieValue('lastName');
  return lastName ? lastName : '';
};

const saveLastName = (name: string) => {
  document.cookie = `lastName=${name};`;
};

export { getLastName, saveLastName };
