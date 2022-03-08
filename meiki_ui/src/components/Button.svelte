<script lang="ts">
    export let label: string
    export let onClick: any = undefined
    export let fullWidth: boolean = false
</script>

<button
    on:click={onClick}
    class="bg-blue-600 hover:bg-blue-500 drop-shadow-card font-bold text-white rounded-xl py-2 px-5"
    class:w-full={fullWidth}
>
    <span>{label}</span>
</button>
