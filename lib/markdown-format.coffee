markdownfmt = require('./markdownfmt.js')

class MarkdownFormat
  constructor: ->
    atom.workspace.eachEditor (editor) =>
      @handleSave(editor)

  destroy: ->
    atom.unsubscribe()

  handleSave: (editor) ->
    buffer = editor.getBuffer()
    atom.subscribe buffer, 'will-be-saved', =>
      if editor.getGrammar().scopeName is 'source.gfm'
        buffer.setTextViaDiff(markdownfmt.ProcessMarkdown(buffer.getText()))

module.exports =
  activate: ->
    @markdownFormat = new MarkdownFormat()

  deactivate: ->
    @markdownFormat.destroy()
