/* eslint-disable no-restricted-globals */
import "./wasm_exec.js";

// eslint-disable-next-line no-undef
const go = new Go();
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then(
  (result) => {
    go.run(result.instance);
    console.log("go running");
  }
);

self.addEventListener("message", (evt) => {
  console.log(evt);

  const { imgArr, colorDiff, minLineLength } = evt.data;
  // eslint-disable-next-line no-undef
  convertToCSS(imgArr, colorDiff, minLineLength, (h) => {
    postMessage(h);
  });
});
