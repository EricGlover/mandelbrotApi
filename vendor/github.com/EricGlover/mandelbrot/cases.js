//note z = [real, imaginary]
const tests = [
  //basic values of z that are in the set
  {
    z: [0, 0],
    answer: true,
    failed: true
  },
  {
    z: [-1, 0],
    answer: true,
    failed: true
  },
  {
    z: [0, 1],
    answer: true,
    failed: true
  },
  {
    z: [-2, 0],
    answer: true,
    failed: true
  },
  //basic values of z that are not in the set, the corners
  {
    z: [-2, 1],
    answer: false,
    failed: true
  },
  {
    z: [1, 1],
    answer: false,
    failed: true
  },
  {
    z: [1, -1],
    answer: false,
    failed: true
  },
  {
    z: [-2, -1],
    answer: false,
    failed: true
  },
  //testing floats
  {
    z: [-0.2, -0.1],
    answer: true,
    failed: true
  },
  {
    z: [0, 0.1],
    answer: true,
    failed: true
  },
  {
    z: [-1, 0.0001],
    answer: true,
    failed: true
  },
  {
    z: [-0.2, 0.2],
    answer: true,
    failed: true
  },
  //zooming in and starting to look for detail
  {
    z: [-0.75, 0],
    answer: true,
    failed: true
  },
  {
    z: [-0.784, 0.344],
    answer: false,
    failed: true
  },
  {
    z: [-0.784, 0.275],
    answer: false,
    failed: true
  },
  {
    z: [-0.767, 0.189],
    answer: false,
    failed: true
  },
  {
    z: [-0.802, 0.12],
    answer: true,
    failed: true
  }
];

