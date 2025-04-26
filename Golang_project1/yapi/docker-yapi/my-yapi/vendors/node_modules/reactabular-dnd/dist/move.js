'use strict';

Object.defineProperty(exports, "__esModule", {
  value: true
});
exports.move = exports.moveRows = exports.moveLabels = exports.moveChildrenLabels = undefined;

var _findIndex4 = require('lodash/findIndex');

var _findIndex5 = _interopRequireDefault(_findIndex4);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

function _defineProperty(obj, key, value) { if (key in obj) { Object.defineProperty(obj, key, { value: value, enumerable: true, configurable: true, writable: true }); } else { obj[key] = value; } return obj; }

function moveChildrenLabels(columns, _ref) {
  var sourceLabel = _ref.sourceLabel,
      targetLabel = _ref.targetLabel;

  var sourceIndex = (0, _findIndex5.default)(columns, function (column) {
    return (0, _findIndex5.default)(column.children, { header: { label: sourceLabel } }) >= 0;
  });

  if (sourceIndex < 0) {
    return null;
  }

  var targetIndex = (0, _findIndex5.default)(columns, function (column) {
    return (0, _findIndex5.default)(column.children, { header: { label: targetLabel } }) >= 0;
  });

  if (targetIndex < 0) {
    return null;
  }

  // Allow drag and drop only within the same column
  if (sourceIndex !== targetIndex) {
    return null;
  }

  var movedChildren = moveLabels(columns[sourceIndex].children, {
    sourceLabel: sourceLabel, targetLabel: targetLabel
  });

  if (!movedChildren) {
    return null;
  }

  return {
    target: sourceIndex,
    columns: movedChildren.columns
  };
}

function moveLabels(columns, _ref2) {
  var sourceLabel = _ref2.sourceLabel,
      targetLabel = _ref2.targetLabel;

  if (!columns) {
    throw new Error('dnd.moveLabels - Missing columns!');
  }

  var sourceIndex = (0, _findIndex5.default)(columns, { header: { label: sourceLabel } });

  if (sourceIndex < 0) {
    return null;
  }

  var targetIndex = (0, _findIndex5.default)(columns, { header: { label: targetLabel } });

  if (targetIndex < 0) {
    return null;
  }

  var movedColumns = move(columns, sourceIndex, targetIndex);

  return {
    source: movedColumns[sourceIndex],
    target: movedColumns[targetIndex],
    columns: movedColumns
  };
}

var moveRows = function moveRows() {
  var _ref3 = arguments.length > 0 && arguments[0] !== undefined ? arguments[0] : {},
      sourceRowId = _ref3.sourceRowId,
      targetRowId = _ref3.targetRowId,
      _ref3$idField = _ref3.idField,
      idField = _ref3$idField === undefined ? 'id' : _ref3$idField;

  return function (rows) {
    var sourceIndex = (0, _findIndex5.default)(rows, _defineProperty({}, idField, sourceRowId));

    if (sourceIndex < 0) {
      return null;
    }

    var targetIndex = (0, _findIndex5.default)(rows, _defineProperty({}, idField, targetRowId));

    if (targetIndex < 0) {
      return null;
    }

    return move(rows, sourceIndex, targetIndex);
  };
};

function move(data, sourceIndex, targetIndex) {
  // Idea
  // a, b, c, d, e -> move(b, d) -> a, c, d, b, e
  // a, b, c, d, e -> move(d, a) -> d, a, b, c, e
  // a, b, c, d, e -> move(a, d) -> b, c, d, a, e
  var sourceItem = data[sourceIndex];

  // 1. detach - a, c, d, e - a, b, c, e, - b, c, d, e
  var ret = data.slice(0, sourceIndex).concat(data.slice(sourceIndex + 1));

  // 2. attach - a, c, d, b, e - d, a, b, c, e - b, c, d, a, e
  return ret.slice(0, targetIndex).concat([sourceItem]).concat(ret.slice(targetIndex));
}

exports.moveChildrenLabels = moveChildrenLabels;
exports.moveLabels = moveLabels;
exports.moveRows = moveRows;
exports.move = move;