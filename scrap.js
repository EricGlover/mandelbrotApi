const json = [
  {
    real: 0,
    imaginary: 0
  },
  {
    real: 0,
    imaginary: 0
  },
  {
    real: 0,
    imaginary: 0
  }
];
const data = [];
for (let i = 0; i < 100; i++) {
  data.push({ real: i * 0.1, imaginary: i * 0.1 });
}

console.log(JSON.stringify(json));
