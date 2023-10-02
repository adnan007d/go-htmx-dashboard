import autoprefixer from "autoprefixer";
import tailwindcss from "tailwindcss";
import postcss from "postcss";
import fs from "fs";
import cssnano from "cssnano";
import esbuild from "esbuild";

// Postcss
fs.readFile("static/css/tailwind.css", (_err, css) => {
  const outDir = "dist/css/";

  if (!fs.existsSync(outDir)) {
    fs.mkdirSync(outDir, { recursive: true });
  }

  postcss([autoprefixer, tailwindcss, cssnano])
    .process(css, { from: "static/css/tailwind.css", to: "dist/css/style.css" })
    .then((result) => {
      fs.writeFile("dist/css/style.css", result.css, () => true);
      if (result.map) {
        fs.writeFile(
          "dist/css/style.css.map",
          result.map.toString(),
          () => true
        );
      }
    });
});

await esbuild.build({
  entryPoints: ["static/js/*.js"],
  bundle: false,
  outdir: "dist/js",
});
