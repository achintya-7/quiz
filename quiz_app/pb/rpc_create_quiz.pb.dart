///
//  Generated code. Do not modify.
//  source: rpc_create_quiz.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:core' as $core;

import 'package:protobuf/protobuf.dart' as $pb;

import 'quiz.pb.dart' as $1;

class CreateQuizRequest extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'CreateQuizRequest', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'pb'), createEmptyInstance: create)
    ..aOS(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'name')
    ..aOS(2, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'description')
    ..aOS(3, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'createdBy')
    ..hasRequiredFields = false
  ;

  CreateQuizRequest._() : super();
  factory CreateQuizRequest({
    $core.String? name,
    $core.String? description,
    $core.String? createdBy,
  }) {
    final _result = create();
    if (name != null) {
      _result.name = name;
    }
    if (description != null) {
      _result.description = description;
    }
    if (createdBy != null) {
      _result.createdBy = createdBy;
    }
    return _result;
  }
  factory CreateQuizRequest.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CreateQuizRequest.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  CreateQuizRequest clone() => CreateQuizRequest()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  CreateQuizRequest copyWith(void Function(CreateQuizRequest) updates) => super.copyWith((message) => updates(message as CreateQuizRequest)) as CreateQuizRequest; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static CreateQuizRequest create() => CreateQuizRequest._();
  CreateQuizRequest createEmptyInstance() => create();
  static $pb.PbList<CreateQuizRequest> createRepeated() => $pb.PbList<CreateQuizRequest>();
  @$core.pragma('dart2js:noInline')
  static CreateQuizRequest getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CreateQuizRequest>(create);
  static CreateQuizRequest? _defaultInstance;

  @$pb.TagNumber(1)
  $core.String get name => $_getSZ(0);
  @$pb.TagNumber(1)
  set name($core.String v) { $_setString(0, v); }
  @$pb.TagNumber(1)
  $core.bool hasName() => $_has(0);
  @$pb.TagNumber(1)
  void clearName() => clearField(1);

  @$pb.TagNumber(2)
  $core.String get description => $_getSZ(1);
  @$pb.TagNumber(2)
  set description($core.String v) { $_setString(1, v); }
  @$pb.TagNumber(2)
  $core.bool hasDescription() => $_has(1);
  @$pb.TagNumber(2)
  void clearDescription() => clearField(2);

  @$pb.TagNumber(3)
  $core.String get createdBy => $_getSZ(2);
  @$pb.TagNumber(3)
  set createdBy($core.String v) { $_setString(2, v); }
  @$pb.TagNumber(3)
  $core.bool hasCreatedBy() => $_has(2);
  @$pb.TagNumber(3)
  void clearCreatedBy() => clearField(3);
}

class CreateQuizResponse extends $pb.GeneratedMessage {
  static final $pb.BuilderInfo _i = $pb.BuilderInfo(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'CreateQuizResponse', package: const $pb.PackageName(const $core.bool.fromEnvironment('protobuf.omit_message_names') ? '' : 'pb'), createEmptyInstance: create)
    ..aOM<$1.Quiz>(1, const $core.bool.fromEnvironment('protobuf.omit_field_names') ? '' : 'quiz', subBuilder: $1.Quiz.create)
    ..hasRequiredFields = false
  ;

  CreateQuizResponse._() : super();
  factory CreateQuizResponse({
    $1.Quiz? quiz,
  }) {
    final _result = create();
    if (quiz != null) {
      _result.quiz = quiz;
    }
    return _result;
  }
  factory CreateQuizResponse.fromBuffer($core.List<$core.int> i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromBuffer(i, r);
  factory CreateQuizResponse.fromJson($core.String i, [$pb.ExtensionRegistry r = $pb.ExtensionRegistry.EMPTY]) => create()..mergeFromJson(i, r);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.deepCopy] instead. '
  'Will be removed in next major version')
  CreateQuizResponse clone() => CreateQuizResponse()..mergeFromMessage(this);
  @$core.Deprecated(
  'Using this can add significant overhead to your binary. '
  'Use [GeneratedMessageGenericExtensions.rebuild] instead. '
  'Will be removed in next major version')
  CreateQuizResponse copyWith(void Function(CreateQuizResponse) updates) => super.copyWith((message) => updates(message as CreateQuizResponse)) as CreateQuizResponse; // ignore: deprecated_member_use
  $pb.BuilderInfo get info_ => _i;
  @$core.pragma('dart2js:noInline')
  static CreateQuizResponse create() => CreateQuizResponse._();
  CreateQuizResponse createEmptyInstance() => create();
  static $pb.PbList<CreateQuizResponse> createRepeated() => $pb.PbList<CreateQuizResponse>();
  @$core.pragma('dart2js:noInline')
  static CreateQuizResponse getDefault() => _defaultInstance ??= $pb.GeneratedMessage.$_defaultFor<CreateQuizResponse>(create);
  static CreateQuizResponse? _defaultInstance;

  @$pb.TagNumber(1)
  $1.Quiz get quiz => $_getN(0);
  @$pb.TagNumber(1)
  set quiz($1.Quiz v) { setField(1, v); }
  @$pb.TagNumber(1)
  $core.bool hasQuiz() => $_has(0);
  @$pb.TagNumber(1)
  void clearQuiz() => clearField(1);
  @$pb.TagNumber(1)
  $1.Quiz ensureQuiz() => $_ensure(0);
}

