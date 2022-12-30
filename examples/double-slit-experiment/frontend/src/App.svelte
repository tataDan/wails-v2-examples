<script>
  import { SingleSlitExperiment } from "../wailsjs/go/main/App";
  import { DoubleSlitExperiment } from "../wailsjs/go/main/App";

  const canvasWidth = 500;
  const canvasHeight = 300;

  var doClearCanvas;

  let points = [];
  let singleBtnDisabled = false;
  let doubleBtnDisabled = false;

  function sleep(ms) {
    return new Promise((resolve) => setTimeout(resolve, ms));
  }

  async function singleSlitExperiment() {
    doubleBtnDisabled = true;
    await SingleSlitExperiment().then((result) => (points = result));
    for (let i = 0; i < points.length; i++) {
      if (i === 0) {
        doClearCanvas = true;
      } else {
        doClearCanvas = false;
      }
      draw(1, points[i].X, points[i].Y);
      await sleep(250);
    }
    doubleBtnDisabled = false;
  }

  async function doubleSlitExperiment() {
    singleBtnDisabled = true;
    await DoubleSlitExperiment().then((result) => (points = result));
    for (let i = 0; i < points.length; i++) {
      if (i === 0) {
        doClearCanvas = true;
      } else {
        doClearCanvas = false;
      }
      draw(2, points[i].X, points[i].Y);
      await sleep(250);
    }
    singleBtnDisabled = false;
  }

  async function draw(canvasNum, x2, y2) {
    var canvas;
    if (canvasNum === 1) {
      canvas = document.getElementById("canvas");
    } else if (canvasNum === 2) {
      canvas = document.getElementById("canvas2");
    }
    // @ts-ignore
    if (canvas.getContext) {
      // @ts-ignore
      const ctx = canvas.getContext("2d");
      const x = canvasWidth / 5;
      const y = canvasHeight / 2;
      if (doClearCanvas) {
        ctx.clearRect(301, 0, 218, canvasHeight);
        ctx.clearRect(100, 120, 130, 65);
        await sleep(500);
      }

      // Draw photon source
      ctx.fillRect(20, 135, 75, 30);
      ctx.font = "10px Georgia";
      ctx.fillText("Photon source", 20, 180);

      drawArrow(ctx, 100, 150, 135, 150, 2, "black");
      drawArrow(ctx, 130, 130, 165, 125, 2, "black");
      drawArrow(ctx, 175, 170, 210, 175, 2, "black");

      // Draw circle on first arrow exiting photon source
      ctx.beginPath();
      ctx.arc(114, 150, 4, 0, 2 * Math.PI);
      ctx.stroke();
      ctx.fill();

      // Draw circle on second arrow exiting photon source
      ctx.beginPath();
      ctx.arc(188, 172, 4, 0, 2 * Math.PI);
      ctx.stroke();
      ctx.fill();

      // Draw circle on third arrow exiting photon source
      ctx.beginPath();
      ctx.arc(143, 128, 4, 0, 2 * Math.PI);
      ctx.stroke();
      ctx.fill();

      // Draw screen with 1 or 2 slits (side view)
      if (canvasNum === 1) {
        ctx.beginPath();
        ctx.moveTo(300, 10);
        ctx.lineTo(300, 140);
        ctx.stroke();
        ctx.beginPath();
        ctx.moveTo(300, 150);
        ctx.lineTo(300, 290);
        ctx.stroke();
        ctx.fillText("Single-Split Screen", 195, 270);
        ctx.fillText("(side view)", 195, 280);
      } else if (canvasNum === 2) {
        ctx.beginPath();
        ctx.moveTo(300, 10);
        ctx.lineTo(300, 100);
        ctx.stroke();
        ctx.beginPath();
        ctx.moveTo(300, 110);
        ctx.lineTo(300, 200);
        ctx.stroke();
        ctx.beginPath();
        ctx.moveTo(300, 210);
        ctx.lineTo(300, 290);
        ctx.stroke();
        ctx.fillText("Double-Split Screen", 195, 270);
        ctx.fillText("(side view)", 195, 280);
      }

      // Draw screen with 1 or 2 slits (front view)
      ctx.fillRect(20, 210, 60, 80);
      if (canvasNum === 1) {
        ctx.fillStyle = "white";
        ctx.fillRect(25, 250, 50, 5);
        ctx.fillStyle = "black";
        ctx.fillText("Single-Split Screen", 85, 220);
      } else if (canvasNum === 2) {
        ctx.fillStyle = "white";
        ctx.fillRect(25, 230, 50, 5);
        ctx.fillRect(25, 265, 50, 5);
        ctx.fillStyle = "black";
        ctx.fillText("Double-Split Screen", 85, 220);
      }
      ctx.fillText("(front view)", 85, 230);

      // Draw detection screen
      ctx.beginPath();
      ctx.moveTo(490, 10);
      ctx.lineTo(490, 290);
      ctx.stroke();

      // Draw point on detection screen
      ctx.beginPath();
      ctx.arc(x2, y2, 4, 0, 2 * Math.PI);
      ctx.stroke();
      ctx.fill();
    }
  }

  function drawArrow(ctx, fromx, fromy, tox, toy, arrowWidth, color) {
    var headlen = 10;
    var angle = Math.atan2(toy - fromy, tox - fromx);

    ctx.save();
    ctx.strokeStyle = color;

    //starting path of the arrow from the start square to the end square
    //and drawing the stroke
    ctx.beginPath();
    ctx.moveTo(fromx, fromy);
    ctx.lineTo(tox, toy);
    ctx.lineWidth = arrowWidth;
    ctx.stroke();

    //starting a new path from the head of the arrow to one of the sides of
    //the point
    ctx.beginPath();
    ctx.moveTo(tox, toy);
    ctx.lineTo(
      tox - headlen * Math.cos(angle - Math.PI / 7),
      toy - headlen * Math.sin(angle - Math.PI / 7)
    );

    //path from the side point of the arrow, to the other side point
    ctx.lineTo(
      tox - headlen * Math.cos(angle + Math.PI / 7),
      toy - headlen * Math.sin(angle + Math.PI / 7)
    );

    //path from the side point back to the tip of the arrow, and then
    //again to the opposite side point
    ctx.lineTo(tox, toy);
    ctx.lineTo(
      tox - headlen * Math.cos(angle - Math.PI / 7),
      toy - headlen * Math.sin(angle - Math.PI / 7)
    );

    //draws the paths created above
    ctx.stroke();
    ctx.restore();
  }
</script>

<main>
  &nbsp;
  <div>
    <canvas id="canvas" width="500" height="300" />
    &nbsp;
    <canvas id="canvas2" width="500" height="300" />
  </div>
  <br />
  <div>
    <span>
      <button disabled={singleBtnDisabled} on:click={singleSlitExperiment}
        >Start Single-Slit Experiment</button
      >
    </span>
    <span>
      <button disabled={doubleBtnDisabled} on:click={doubleSlitExperiment}
        >Start Double-Slit Experiment</button
      >
    </span>
  </div>
</main>

<style>
  canvas {
    border: 1px solid black;
  }

  div {
    width: 100%;
  }

  span {
    display: inline-block;
    width: 47%;
    text-align: center;
  }
</style>
