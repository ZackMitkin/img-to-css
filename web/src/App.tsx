import { Container, Input, Typography } from "@material-ui/core";
import Grid from "@material-ui/core/Grid/Grid";
import { ChangeEventHandler, useState } from "react";
import Image from "material-ui-image";

declare const convertToCSS: (
  buffer: ArrayBuffer,
  callback: (html: string) => any
) => any;

function App() {
  const [imageFile, setImageFile] = useState<File>();
  const [imageFileUrl, setImageFileUrl] = useState<string>();

  const [html, setHtml] = useState<string>();

  const onSelectedFile: ChangeEventHandler<HTMLInputElement> = async (item) => {
    const file = item.target.files?.[0];
    if (file) {
      setImageFile(file);
      setImageFileUrl(URL.createObjectURL(file));

      const reader = new FileReader();
      reader.onload = function () {
        const buffer = this.result as ArrayBuffer;
        const array = new Uint8Array(buffer);
        convertToCSS(array, (h) => {
          setHtml(h);
        });
      };
      reader.readAsArrayBuffer(file);
    }
  };

  return (
    <Container>
      <Grid container>
        <Grid item xs={6}>
          <Typography variant={"subtitle1"}>
            Select an image file to get started!
          </Typography>
          <Input type={"file"} onChange={onSelectedFile} />
          {imageFileUrl && (
            <Image src={imageFileUrl} imageStyle={{ objectFit: "contain" }} />
          )}
        </Grid>
        <Grid item xs={6}>
          <iframe srcDoc={html} title="image" width={"100%"} height={"100%"} style={{border: 'none'}}/>
        </Grid>
      </Grid>
    </Container>
  );
}

export default App;
