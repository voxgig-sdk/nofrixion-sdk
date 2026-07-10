<?php
declare(strict_types=1);

// Nofrixion SDK utility registration

require_once __DIR__ . '/../core/UtilityType.php';
require_once __DIR__ . '/Clean.php';
require_once __DIR__ . '/Done.php';
require_once __DIR__ . '/MakeError.php';
require_once __DIR__ . '/FeatureAdd.php';
require_once __DIR__ . '/FeatureHook.php';
require_once __DIR__ . '/FeatureInit.php';
require_once __DIR__ . '/Fetcher.php';
require_once __DIR__ . '/MakeFetchDef.php';
require_once __DIR__ . '/MakeContext.php';
require_once __DIR__ . '/MakeOptions.php';
require_once __DIR__ . '/MakeRequest.php';
require_once __DIR__ . '/MakeResponse.php';
require_once __DIR__ . '/MakeResult.php';
require_once __DIR__ . '/MakePoint.php';
require_once __DIR__ . '/MakeSpec.php';
require_once __DIR__ . '/MakeUrl.php';
require_once __DIR__ . '/Param.php';
require_once __DIR__ . '/PrepareAuth.php';
require_once __DIR__ . '/PrepareBody.php';
require_once __DIR__ . '/PrepareHeaders.php';
require_once __DIR__ . '/PrepareMethod.php';
require_once __DIR__ . '/PrepareParams.php';
require_once __DIR__ . '/PreparePath.php';
require_once __DIR__ . '/PrepareQuery.php';
require_once __DIR__ . '/ResultBasic.php';
require_once __DIR__ . '/ResultBody.php';
require_once __DIR__ . '/ResultHeaders.php';
require_once __DIR__ . '/TransformRequest.php';
require_once __DIR__ . '/TransformResponse.php';

NofrixionUtility::setRegistrar(function (NofrixionUtility $u): void {
    $u->clean = [NofrixionClean::class, 'call'];
    $u->done = [NofrixionDone::class, 'call'];
    $u->make_error = [NofrixionMakeError::class, 'call'];
    $u->feature_add = [NofrixionFeatureAdd::class, 'call'];
    $u->feature_hook = [NofrixionFeatureHook::class, 'call'];
    $u->feature_init = [NofrixionFeatureInit::class, 'call'];
    $u->fetcher = [NofrixionFetcher::class, 'call'];
    $u->make_fetch_def = [NofrixionMakeFetchDef::class, 'call'];
    $u->make_context = [NofrixionMakeContext::class, 'call'];
    $u->make_options = [NofrixionMakeOptions::class, 'call'];
    $u->make_request = [NofrixionMakeRequest::class, 'call'];
    $u->make_response = [NofrixionMakeResponse::class, 'call'];
    $u->make_result = [NofrixionMakeResult::class, 'call'];
    $u->make_point = [NofrixionMakePoint::class, 'call'];
    $u->make_spec = [NofrixionMakeSpec::class, 'call'];
    $u->make_url = [NofrixionMakeUrl::class, 'call'];
    $u->param = [NofrixionParam::class, 'call'];
    $u->prepare_auth = [NofrixionPrepareAuth::class, 'call'];
    $u->prepare_body = [NofrixionPrepareBody::class, 'call'];
    $u->prepare_headers = [NofrixionPrepareHeaders::class, 'call'];
    $u->prepare_method = [NofrixionPrepareMethod::class, 'call'];
    $u->prepare_params = [NofrixionPrepareParams::class, 'call'];
    $u->prepare_path = [NofrixionPreparePath::class, 'call'];
    $u->prepare_query = [NofrixionPrepareQuery::class, 'call'];
    $u->result_basic = [NofrixionResultBasic::class, 'call'];
    $u->result_body = [NofrixionResultBody::class, 'call'];
    $u->result_headers = [NofrixionResultHeaders::class, 'call'];
    $u->transform_request = [NofrixionTransformRequest::class, 'call'];
    $u->transform_response = [NofrixionTransformResponse::class, 'call'];
});
