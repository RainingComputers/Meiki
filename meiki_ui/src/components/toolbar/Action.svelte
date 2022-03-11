<script lang="ts">
    export let label: string
    export let onClick: any = undefined
</script>

<button
    on:click={onClick}
    class="bg-blue-600 hover:bg-blue-500 drop-shadow-card text-white rounded-xl py-1 px-5"
>
    <span>{label}</span>
</button>
