<script>
    export let fullWidth = false
</script>

<span
    class=" text-green-500 font-bold bg-green-50 p-4 rounded-lg border-2 border-green-400 drop-shadow-card"
    class:w-full={fullWidth}
>
    <slot />
</span>
