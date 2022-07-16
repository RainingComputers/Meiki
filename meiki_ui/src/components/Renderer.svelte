<script lang="ts">
    // TODO style this if needed

    import Markdoc from "@markdoc/markdoc"

    export let text: string = ""

    $: ast = Markdoc.parse(text)
    $: content = Markdoc.transform(ast)
</script>

<p data-cy="renderer" class="p-2 overflow-scroll h-full prose max-w-full">
    {@html Markdoc.renderers.html(content)}
</p>
