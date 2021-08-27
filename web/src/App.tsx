import {
  Box,
  CircularProgress,
  Container,
  Input,
  Slider,
  Typography,
} from "@material-ui/core";
import Grid from "@material-ui/core/Grid/Grid";
import { ChangeEventHandler, useEffect, useState } from "react";
import Image from "material-ui-image";

const worker = new Worker("./worker.js", { type: "module" });
function App() {
  const [imgArr, setImgArr] = useState<Uint8Array>();
  const [imageFileUrl, setImageFileUrl] = useState<string>();
  const [html, setHtml] = useState<string>();
  const [loading, setLoading] = useState(false);
  const [colorDiff, setColorDiff] = useState(40);
  const [minLineLength, setMinLineLength] = useState(10);

  useEffect(() => {
    console.log("start");

    worker.onmessage = (evt) => {
      setHtml(evt.data);
      setLoading(false);
    };
    return () => worker.terminate();
  }, []);

  function process() {
    setLoading(true)
    worker.postMessage({ imgArr, colorDiff, minLineLength });
  }

  const onSelectedFile: ChangeEventHandler<HTMLInputElement> = async (item) => {
    const file = item.target.files?.[0];
    if (file) {
      setImageFileUrl(URL.createObjectURL(file));
      const reader = new FileReader();
      reader.onload = async function () {
        const buffer = this.result as ArrayBuffer;
        const array = new Uint8Array(buffer);
        setImgArr(array);
        worker.postMessage({ imgArr: array, colorDiff, minLineLength });
        setLoading(true);
      };
      reader.readAsArrayBuffer(file);
    }
  };

  return (
    <Container>
      <Grid container>
        <Grid item xs={6}>
          {imageFileUrl && (
            <Image src={imageFileUrl} imageStyle={{ objectFit: "contain" }} />
          )}
        </Grid>
        <Grid item xs={6}>
          <Box
            display={"flex"}
            justifyContent={"center"}
            alignItems={"center"}
            height={"100%"}
          >
            {loading ? (
              <CircularProgress />
            ) : (
              <iframe
                srcDoc={html}
                title="image"
                width={"100%"}
                height={"100%"}
                scrolling="no"
                style={{ border: "none", overflow: "clip" }}
              />
            )}
          </Box>
        </Grid>
      </Grid>
      <Grid container justifyContent={"center"}>
        <Grid item xs={4}>
          <Box p={1}>
            <Input type={"file"} onChange={onSelectedFile} />
          </Box>
          <Box p={1}>
            <Typography>Color Difference: {colorDiff}</Typography>
            <Slider
              value={colorDiff}
              onChange={(_, v) => setColorDiff(v as number)}
              onChangeCommitted={process}
            />
          </Box>
          <Box p={1}>
            <Typography>
              Minimal Polygon Line Length (px): {minLineLength}
            </Typography>
            <Slider
              value={minLineLength}
              onChange={(_, v) => setMinLineLength(v as number)}
              onChangeCommitted={process}
            />
          </Box>
        </Grid>
      </Grid>
    </Container>
  );
}

export default App;
