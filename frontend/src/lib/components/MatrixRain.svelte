<script lang="ts">
	import { onMount } from 'svelte';

	let canvas: HTMLCanvasElement;

	onMount(() => {
		const ctx = canvas.getContext('2d')!;

		function resize() {
			canvas.width = window.innerWidth;
			canvas.height = window.innerHeight;
			initParticles();
		}

		const symbols = ['~', '*', '.', '+', 'o', ':', '^'];
		const colors = ['#7bc89b', '#a8ddb8', '#5db87e', '#b8e6c8', '#9ccfd8'];

		interface Particle {
			x: number;
			y: number;
			speed: number;
			size: number;
			color: string;
			symbol: string;
			opacity: number;
		}

		let particles: Particle[] = [];

		function createParticle(): Particle {
			return {
				x: Math.random() * canvas.width,
				y: Math.random() * canvas.height,
				speed: 0.2 + Math.random() * 0.5,
				size: 32 + Math.random() * 40,
				color: colors[Math.floor(Math.random() * colors.length)],
				symbol: symbols[Math.floor(Math.random() * symbols.length)],
				opacity: 0.15 + Math.random() * 0.25
			};
		}

		function initParticles() {
			const count = Math.floor((canvas.width * canvas.height) / 18000);
			particles = Array.from({ length: count }, createParticle);
		}

		resize();
		window.addEventListener('resize', resize);

		let lastTime = 0;

		function draw(time: number) {
			const delta = lastTime ? time - lastTime : 16;
			lastTime = time;

			ctx.clearRect(0, 0, canvas.width, canvas.height);

			for (const p of particles) {
				ctx.globalAlpha = p.opacity;
				ctx.fillStyle = p.color;
				ctx.font = `${p.size}px monospace`;
				ctx.fillText(p.symbol, p.x, p.y);

				p.y -= p.speed * delta * 0.03;
				if (p.y < -20) {
					p.y = canvas.height + 20;
					p.x = Math.random() * canvas.width;
				}
			}
			ctx.globalAlpha = 1;

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
