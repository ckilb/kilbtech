class MatrixRevealAnimator {
    constructor(imgElement) {
        this.img = imgElement;
        this.wrapper = this.img.parentElement;
        this.canvas = document.createElement('canvas');
        this.ctx = this.canvas.getContext('2d', { willReadFrequently: true });

        this.originalData = null;
        this.w = 0;
        this.h = 0;

        this.startTime = null;
        this.duration = 10000; // 10 seconds total duration to allow for a slower start
        this.isRunning = false;

        this.chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789@#$%^&*()_+-=~[]{}|;:,.<>?';
        this.baseFontSize = 14;

        this.charMap = new Array(8000).fill('');
        for(let i=0; i<this.charMap.length; i++) this.charMap[i] = this.randomChar();

        this.init();
    }

    init() {
        if (this.img.complete) {
            this.start();
        } else {
            this.img.addEventListener('load', () => this.start());
        }
    }

    randomChar() {
        return this.chars[Math.floor(Math.random() * this.chars.length)];
    }

    start() {
        this.w = this.img.width || this.img.clientWidth || 400;
        this.h = this.img.height || this.img.clientHeight || 400;

        if (this.w === 0 || this.h === 0) {
            setTimeout(() => this.start(), 100);
            return;
        }

        this.canvas.width = this.w;
        this.canvas.height = this.h;
        this.img.style.opacity = '0';

        this.canvas.className = this.img.className;
        this.canvas.id = 'hero-canvas-matrix';
        this.canvas.style.position = 'absolute';
        this.canvas.style.top = '0';
        this.canvas.style.left = '0';
        this.canvas.style.zIndex = '15';

        this.wrapper.appendChild(this.canvas);

        // Extract pixel data
        let offCanvas = document.createElement('canvas');
        offCanvas.width = this.w;
        offCanvas.height = this.h;
        let offCtx = offCanvas.getContext('2d');
        offCtx.drawImage(this.img, 0, 0, this.w, this.h);
        try {
            this.originalData = offCtx.getImageData(0, 0, this.w, this.h);
        } catch (e) {
            // CORS fallback for local testing
            console.error(e);
            this.img.style.opacity = '1';
            if(this.canvas.parentNode) this.canvas.parentNode.removeChild(this.canvas);
            return;
        }

        this.isRunning = true;
        this.startTime = performance.now();
        requestAnimationFrame((t) => this.drawLoop(t));
    }

    drawLoop(currentTime) {
        if (!this.isRunning) return;

        let elapsed = currentTime - this.startTime;

        if (elapsed > this.duration) {
            this.finish();
            return;
        }

        this.ctx.clearRect(0, 0, this.w, this.h);

        // Update a random subset of charMap to create an ongoing, persistent Matrix "flicker"
        for(let i=0; i<100; i++) {
            this.charMap[Math.floor(Math.random() * this.charMap.length)] = this.randomChar();
        }

        // Sequence Timing
        let t1_color_blend_start = 2000; // Pause for 2s on pure green text
        let t1_color_blend_end = 4000;

        let t2_shrink_start = 2700; // Start shrinking later so the starting silhouette holds
        let t2_shrink_end = 10000;   // Shrink all the way to the end for a heavily drawn-out, slow ending

        let t3_real_fade_start = 6000; // Real image fades in over a long 4s window for a buttery smooth crossfade
        let t3_real_fade_end = 10000;

        // Color Blend Logic: 0 = Pure Green, 1 = True Image Color
        let colorBlend = 0;
        if (elapsed > t1_color_blend_start) {
            colorBlend = (elapsed - t1_color_blend_start) / (t1_color_blend_end - t1_color_blend_start);
            if(colorBlend > 1) colorBlend = 1;
        }

        // Shrink Logic: Shrink font size over time to create more dense "pixels"
        let currentFontSize = this.baseFontSize;
        if (elapsed > t2_shrink_start) {
            let shrinkProgress = (elapsed - t2_shrink_start) / (t2_shrink_end - t2_shrink_start);
            if (shrinkProgress > 1) shrinkProgress = 1;

            // easeOutQuart curve makes it start fast and slow down significantly at the end
            let easeOut = 1 - Math.pow(1 - shrinkProgress, 4);

            // Stop shrinking at 3px. If we go smaller, we attempt to draw 100k+ rectangles per frame
            // which causes intense browser lag and results in a "keyframe" or stuttering effect.
            currentFontSize = this.baseFontSize - (easeOut * (this.baseFontSize - 3));
        }

        let realFadeProgress = 0;
        if (elapsed > t3_real_fade_start) {
            realFadeProgress = (elapsed - t3_real_fade_start) / (t3_real_fade_end - t3_real_fade_start);
            if (realFadeProgress > 1) realFadeProgress = 1;
        }
        // Smoothstep easing for a much smoother crossfade between pixel canvas and real image
        let realFade = realFadeProgress * realFadeProgress * (3 - 2 * realFadeProgress);
        this.img.style.opacity = realFade;

        // Only draw the canvas if the real image hasn't fully taken over
        if (realFade < 1) {
            this.ctx.font = 'bold ' + Math.max(1, currentFontSize) + 'px monospace';

            // Symmetrical grid expansion from center
            let halfCols = Math.ceil((this.w / currentFontSize) / 2);
            let halfRows = Math.ceil((this.h / currentFontSize) / 2);
            let cx = this.w / 2;
            let cy = this.h / 2;

            for(let y = -halfRows; y <= halfRows; y++) {
                for(let x = -halfCols; x <= halfCols; x++) {

                    let px = Math.floor(cx + x * currentFontSize);
                    let py = Math.floor(cy + y * currentFontSize);

                    if (px < 0 || px >= this.w || py < 0 || py >= this.h) continue;

                    let idx = (py * this.w + px) * 4;

                    if(idx >= 0 && idx < this.originalData.data.length) {
                        let r = this.originalData.data[idx];
                        let g = this.originalData.data[idx+1];
                        let b = this.originalData.data[idx+2];
                        let alpha = this.originalData.data[idx+3];

                        // Skip fully transparent pixels (like the pure white/black background mask if any)
                        if (alpha < 10) continue;

                        let brightness = (r+g+b)/3;

                        // REMOVED BLACK LETTERS: Only draw if the source pixel has some brightness
                        // This prevents dark/black areas from rendering text characters
                        if (brightness < 15) continue;

                        // Color Shift:
                        // Start as Green (0, brightness, 0)
                        // Transition to Real (r, g, b)
                        let drawR = Math.floor(r * colorBlend);
                        let drawG = Math.floor(brightness * (1 - colorBlend) + g * colorBlend);
                        let drawB = Math.floor(b * colorBlend);

                        // Fade canvas out as real image fades in
                        let canvasAlpha = 1 - realFade;

                        this.ctx.fillStyle = `rgba(${drawR}, ${drawG}, ${drawB}, ${canvasAlpha})`;
                        // Draw random character if text is large, otherwise just a tiny square (pixel)
                        if (currentFontSize > 3) {
                            let charList = '01';
                            let fixedChar;
                            if (currentFontSize > 6) {
                                // Map coordinates to a persistent random char that changes naturally over time via charMap updates
                                let charIdx = Math.abs(x * 137 + y * 19) % this.charMap.length;
                                fixedChar = this.charMap[charIdx];
                            } else {
                                fixedChar = charList[Math.abs(x*y) % charList.length];
                            }

                            // Center alignment for chars
                            this.ctx.textAlign = 'center';
                            this.ctx.textBaseline = 'middle';
                            this.ctx.fillText(fixedChar, px, py);
                        } else {
                            // At very small sizes, text rendering is incredibly expensive and illegible,
                            // so we switch to drawing rectangles (pixels) for the final transition
                            this.ctx.fillRect(px - currentFontSize/2, py - currentFontSize/2, currentFontSize, currentFontSize);
                        }
                    }
                }
            }
        }

        requestAnimationFrame((t) => this.drawLoop(t));
    }

    finish() {
        this.isRunning = false;
        this.img.style.opacity = '1';
        this.canvas.style.display = 'none'; // Instant hide since real img is 100% opacity
        setTimeout(() => {
            if (this.canvas && this.canvas.parentNode) {
                this.canvas.parentNode.removeChild(this.canvas);
            }
        }, 100);
    }
}

document.addEventListener("DOMContentLoaded", () => {
    setTimeout(() => {
        const heroImg = document.getElementById('hero-img-matrix');
        if (heroImg) {
            new MatrixRevealAnimator(heroImg);
        }
    }, 100);
});
