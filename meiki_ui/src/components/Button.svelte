<script lang="ts">
    export let label: string
    export let onClick: any = undefined
    export let fullWidth: boolean = false
    export let type: string = "primary" //TODO: Make enum

    const classPrimary =
        type == "primary" ? "bg-blue-600 hover:bg-blue-500" : ""
    const classSecondary =
        type == "secondary" ? "bg-slate-600 hover:bg-slate-500" : ""
</script>

<button
    on:click={onClick}
    class="drop-shadow-card font-bold text-white rounded-xl py-2 px-5 {classPrimary} {classSecondary}"
    class:w-full={fullWidth}
>
    <span>{label}</span>
</button>
