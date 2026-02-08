<script lang="ts">
	import { onMount } from 'svelte';

	let canvas: HTMLCanvasElement;

	onMount(() => {
		const ctx = canvas.getContext('2d')!;

		function resize() {
			canvas.width = window.innerWidth;
			canvas.height = window.innerHeight;
			initColumns();
		}

		const chars = 'アイウエオカキクケコサシスセソタチツテトナニヌネノハヒフヘホマミムメモヤユヨラリルレロワヲン0123456789';
		const fontSize = 14;

		let columns = 0;
		let drops: number[] = [];
		let speeds: number[] = [];

		// speeds in rows/ms: fastest ~3s to cross, slowest ~10s
		function randomSpeed() {
			return 0.008 + Math.random() * 0.018;
		}

		function initColumns() {
			const newCols = Math.floor(canvas.width / fontSize);
			if (newCols !== columns) {
				columns = newCols;
				drops = Array(columns).fill(0).map(() => Math.random() * -50);
				speeds = Array(columns).fill(0).map(randomSpeed);
			}
		}

		resize();
		window.addEventListener('resize', resize);

		let lastTime = 0;

		function draw(time: number) {
			const delta = lastTime ? time - lastTime : 16;
			lastTime = time;

			ctx.fillStyle = 'rgba(0, 0, 0, 0.04)';
			ctx.fillRect(0, 0, canvas.width, canvas.height);

			ctx.fillStyle = '#00ff41';
			ctx.font = `${fontSize}px monospace`;

			for (let i = 0; i < drops.length; i++) {
				const char = chars[Math.floor(Math.random() * chars.length)];
				const x = i * fontSize;
				const y = drops[i] * fontSize;

				ctx.globalAlpha = 0.3 + Math.random() * 0.4;
				ctx.fillText(char, x, y);
				ctx.globalAlpha = 1;

				if (y > canvas.height && Math.random() > 0.98) {
					drops[i] = 0;
					speeds[i] = randomSpeed();
				}
				drops[i] += speeds[i] * delta;
			}

			requestAnimationFrame(draw);
		}

		requestAnimationFrame(draw);

		return () => {
			window.removeEventListener('resize', resize);
		};
	});
</script>

<canvas bind:this={canvas}></canvas>

<style>
	canvas {
		position: fixed;
		top: 0;
		left: 0;
		width: 100%;
		height: 100%;
		z-index: -1;
		pointer-events: none;
	}
</style>
